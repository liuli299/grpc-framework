package controller

import (
	"time"
	"fmt"
	"golang.org/x/net/context"
	pb "github.com/liuli299/grpc-framework/proto/index"
	)

func (s *Server) Test(ctx context.Context, in *pb.IndexRequest) (*pb.IndexReply, error) {
	fmt.Printf("recv a request:%s\n", time.Now())
	return &pb.IndexReply{Message: "Hello " + in.Name}, nil
}

func (s *Server) Alie(ctx context.Context, in *pb.IndexRequest) (*pb.IndexReply, error) {
	fmt.Printf("recv a request:%s\n", time.Now())
	return &pb.IndexReply{Message: "Alie " + in.Name}, nil
}
