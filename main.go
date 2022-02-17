package main

import (
	"fmt"
	"io"
	"net"
	"sort"
	"sync"
	"sync/atomic"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	pb "github.com/hicqu/phybr/pkg/phybr"
)

type X struct {
	instances   int
	offset      *uint32
	received    *sync.WaitGroup
	regionMetas [][]*pb.RegionMeta
}

func newX(instances int) X {
	var offset = new(uint32)
	var received = new(sync.WaitGroup)
	received.Add(instances)
	var regionMetas = make([][]*pb.RegionMeta, instances)
	return X{instances, offset, received, regionMetas}
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
	fmt.Printf("receive reports from one instance finish\n")
	// Do somethine here...
	x.received.Done()
	return
}

func main() {
	address := "127.0.0.1:3379"
	var x = newX(2)

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
	selectReplicas(x.regionMetas)

	fmt.Printf("stopping...")
	s.Stop()
}

func selectReplicas(regionMetas [][]*pb.RegionMeta) [][]*pb.RegionRecover {
	type M struct {
		*pb.RegionMeta
		offset int
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
	// Reverse sort replicas by applied index.
	for rid := range regions {
		var x = regions[rid]
		sort.Slice(x, func(i, j int) bool { return x[i].AppliedIndex > x[j].AppliedIndex })
	}

	return nil
}
