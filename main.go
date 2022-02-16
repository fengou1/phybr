package main

import (
	"fmt"
	"io"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	pb "github.com/hicqu/phybr/pkg/phybr"
)

type srv struct {
	done        chan<- struct{}
	regionMetas []*pb.RegionMeta
}

func newSrv(done chan<- struct{}) srv {
	return srv{done: done, regionMetas: make([]*pb.RegionMeta, 0)}
}

func (s srv) RecoverRegions(stream pb.Phybr_RecoverRegionsServer) (err error) {
	for {
		var meta *pb.RegionMeta
		if meta, err = stream.Recv(); err != nil {
			if err == io.EOF {
				break
			}
			fmt.Errorf("receive fail: %v\n", err)
			return err
		}
		s.regionMetas = append(s.regionMetas, meta)
	}
	// Do somethine here...
	s.done <- struct{}{}
	return
}

func main() {
	address := "127.0.0.7:3379"
	instances := 2

	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Errorf("listen %s fail: %v\n", address, err)
		return
	}
	fmt.Printf("listen %s success, waiting for %d instances\n", address, instances)

	done := make(chan struct{}, instances)
	s := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time:    time.Duration(10) * time.Second,
			Timeout: time.Duration(3) * time.Second,
		}),
	)
	pb.RegisterPhybrServer(s, newSrv(done))
	s.Serve(listener)

	for i := 0; i < instances; i++ {
		<-done
	}
	s.Stop()
}
