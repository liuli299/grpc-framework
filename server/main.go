/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

//go:generate protoc -I ../proto/index --go_out=plugins=grpc:../proto/index ../proto/index/index.proto

package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/liuli299/grpc-framework/proto/index"
	"github.com/liuli299/grpc-framework/server/controller"
)

const (
	port = ":60051"
)

// server is used to implement helloworld.GreeterServer.

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterIndexServer(s, &controller.Server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
