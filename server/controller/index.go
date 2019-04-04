package controller

import (
	"golang.org/x/net/context"
	pb "github.com/liuli299/grpc-framework/proto/index"
	)

func (s *Server) Post(ctx context.Context, in *pb.IndexRequest) (*pb.IndexReply, error) {
	var data = make([]*pb.Data, 0)
	var tmp pb.Data
	tmp.Title = "ffff"
	tmp.Img = "dasdasd"
	tmp.Url = in.Uid
	data = append(data, &tmp)
	return &pb.IndexReply{Code:1, Data:data}, nil
}

func (s *Server) Get(ctx context.Context, in *pb.IndexRequest) (*pb.IndexReply, error) {
	var data = make([]*pb.Data, 0)
	var tmp pb.Data
	tmp.Title = "aaaa"
	tmp.Img = "22222"
	tmp.Url = in.Uid
	data = append(data, &tmp)
	return &pb.IndexReply{Code:1, Data:data}, nil
}
