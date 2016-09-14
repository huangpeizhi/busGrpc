package main

import (
	"net"
	"sync/atomic"

	"github.com/go-playground/log"

	pb "busGrpc/hservice"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct {
	sCnt int64
}

func (s *server) Send(ctx context.Context, in *pb.MessageRequest) (*pb.MessageReply, error) {
	atomic.AddInt64(&s.sCnt, 1)

	log.WithFields(log.F("func", "server.Send"), log.F("sCnt", s.sCnt)).Info(in.String())
	return &pb.MessageReply{Ok: true}, nil
}

func gRpcRun() {
	lis, err := net.Listen("tcp", opts.Hook.Addr)
	if err != nil {
		log.WithFields(log.F("func", "gRpcRun")).Fatal(err.Error())
	}
	s := grpc.NewServer()
	pb.RegisterHookServiceServer(s, &server{sCnt: 0})

	s.Serve(lis)
}
