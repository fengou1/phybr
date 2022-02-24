package main

import (
	"fmt"
	"io"
	"net"
	"sort"
	"sync"
	"sync/atomic"
	"time"

	"github.com/emirpasic/gods/maps/treemap"
	pb "github.com/hicqu/phybr/pkg/phybr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type X struct {
	instances      int
	offset         *uint32
	received       *sync.WaitGroup // Whether all reports are collected or not.
	generated      *sync.WaitGroup // Whether commands are generated or not.
	finished       *sync.WaitGroup // Whether all commands are sent out or not.
	regionMetas    [][]*pb.RegionMeta
	regionRecovers [][]*pb.RegionRecover
}

func newX(instances int) X {
	var offset = new(uint32)
	var received = new(sync.WaitGroup)
	received.Add(instances)
	var generated = new(sync.WaitGroup)
	generated.Add(1)
	var finished = new(sync.WaitGroup)
	finished.Add(instances)
	var regionMetas = make([][]*pb.RegionMeta, instances)
	var regionRecovers = make([][]*pb.RegionRecover, instances)
	return X{instances, offset, received, generated, finished, regionMetas, regionRecovers}
}

func (x X) RecoverRegions(stream pb.Phybr_RecoverRegionsServer) (err error) {
	var offset = atomic.AddUint32(x.offset, 1) - 1
	for {
		var meta *pb.RegionMeta
		if meta, err = stream.Recv(); err == nil {
			x.regionMetas[offset] = append(x.regionMetas[offset], meta)
		} else if err == io.EOF {
			break
		} else {
			fmt.Errorf("receive fail: %v\n", err)
			return err
		}
	}
	fmt.Printf("receive reports from instance %d finish\n", offset)
	x.received.Done()

	x.generated.Wait()
	if err = stream.SendHeader(nil); err != nil {
		fmt.Errorf("send fail: %v\n", err)
		return err
	}
	for _, command := range x.regionRecovers[offset] {
		if err = stream.Send(command); err != nil {
			fmt.Errorf("send fail: %v\n", err)
			return err
		}
	}
	fmt.Printf("all commands are sent to instance %d\n", offset)
	x.finished.Done()
	return
}

func main() {
	address := "127.0.0.1:3379"
	var x = newX(3)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Errorf("listen %s fail: %v\n", address, err)
		return
	}
	fmt.Printf("listen %s success, waiting for %d instances\n", address, x.instances)

	s := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time:    time.Duration(10) * time.Second,
			Timeout: time.Duration(3) * time.Second,
		}),
	)

	pb.RegisterPhybrServer(s, x)
	go s.Serve(listener)

	// Wait for all TiKV instances reporting region metas.
	x.received.Wait()
	selectReplicas(x.regionMetas, x.regionRecovers)
	fmt.Printf("commands are generated\n")
	x.generated.Done()

	x.finished.Wait()
	time.Sleep(time.Second * 3) // Sleep to wait grpc streams get closed.
	fmt.Printf("stopping...")
	s.Stop()
}

func selectReplicas(regionMetas [][]*pb.RegionMeta, commands [][]*pb.RegionRecover) {
	type M struct {
		*pb.RegionMeta
		offset int
	}

	type P struct {
		r uint64
		v uint64
	}

	type T struct {
		k []byte
		r uint64
	}

	// Group region metas by region id.
	var regions = make(map[uint64][]M, 0)
	for offset, v := range regionMetas {
		for _, m := range v {
			if regions[m.RegionId] == nil {
				regions[m.RegionId] = make([]M, 0, 3)
			}
			regions[m.RegionId] = append(regions[m.RegionId], M{m, offset})
		}
	}

	// Reverse sort replicas by applied index, and collect all regions' version.
	var versions = make([]P, 0, len(regions))
	for r, x := range regions {
		sort.Slice(x, func(i, j int) bool { return x[i].AppliedIndex > x[j].AppliedIndex })
		var v = x[0].Version
		versions = append(versions, P{r, v})
	}
	sort.Slice(versions, func(i, j int) bool { return versions[i].v > versions[j].v })

	// Resolve version conflicts.
	var topo = treemap.NewWith(keyCmpInterface)
	for _, p := range versions {
		var sk = prefixStartKey(regions[p.r][0].StartKey)
		var ek = prefixEndKey(regions[p.r][0].EndKey)
		var fk, fv interface{}
		fk, _ = topo.Ceiling(sk)
		if fk != nil && (keyEq(fk.([]byte), sk) || keyCmp(fk.([]byte), ek) < 0) {
			continue
		}
		fk, fv = topo.Floor(sk)
		if fk != nil && keyCmp(fv.(T).k, sk) > 0 {
			continue
		}
		topo.Put(sk, T{ek, p.r})
	}

	// After resolved, all reserved regions shouldn't be tombstone.
	var reserved = make(map[uint64]struct{}, 0)
	var iter = topo.Iterator()
	var prevEndKey = prefixStartKey([]byte{})
	var prevR uint64 = 0
	for iter.Next() {
		v := iter.Value().(T)
		if regions[v.r][0].Tombstone {
			fmt.Printf("reserved region shouldn't be tombstone: %d\n", v.r)
			panic("reserved region shouldn't be tombstone")
		}
		if !keyEq(prevEndKey, iter.Key().([]byte)) {
			fmt.Printf("region %d doesn't conject to region %d\n", prevR, v.r)
			panic("regions should conject to each other")
		}
		prevEndKey = v.k
		prevR = v.r
		reserved[v.r] = struct{}{}
	}

	// Generate recover commands.
	for r, x := range regions {
		if _, ok := reserved[r]; !ok {
			// Generate a tombstone command.
			for _, m := range x {
				cmd := &pb.RegionRecover{RegionId: m.RegionId, Tombstone: true}
				if cmd.Tombstone != m.Tombstone {
					commands[m.offset] = append(commands[m.offset], cmd)
				}
			}
		} else {
			// Generate normal commands.
			var maxTerm uint64 = 0
			for _, m := range x {
				if m.Term > maxTerm {
					maxTerm = m.Term
				}
			}
			for i, m := range x {
				cmd := &pb.RegionRecover{RegionId: m.RegionId, Term: maxTerm}
				cmd.Silence = i > 0
				if cmd.Term != m.Term || cmd.Silence {
					commands[m.offset] = append(commands[m.offset], cmd)
				}
			}
		}
	}
}

func prefixStartKey(key []byte) []byte {
	var sk = make([]byte, 0, len(key)+1)
	sk = append(sk, 'z')
	sk = append(sk, key...)
	return sk
}

func prefixEndKey(key []byte) []byte {
	if len(key) == 0 {
		return []byte{'z' + 1}
	}
	return prefixStartKey(key)
}

func keyEq(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func keyCmp(a, b []byte) int {
	var length = 0
	var chosen = 0
	if len(a) < len(b) {
		length = len(a)
		chosen = -1
	} else if len(a) == len(b) {
		length = len(a)
		chosen = 0
	} else {
		length = len(b)
		chosen = 1
	}
	for i := 0; i < length; i++ {
		if a[i] < b[i] {
			return -1
		} else if a[i] > b[i] {
			return 1
		}
	}
	return chosen
}

func keyCmpInterface(a, b interface{}) int {
	return keyCmp(a.([]byte), b.([]byte))
}
