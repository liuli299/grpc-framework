package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	gw "github.com/liuli299/grpc-framework/proto/index"
)

var (
	proxyAddress = "127.0.0.1:8080"
	grpcAddress = "127.0.0.1:9090"
	endpoint = flag.String("grpc_server", grpcAddress, "endpoint of IndexServer")
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterIndexHandlerFromEndpoint(ctx, mux, *endpoint, opts)
	if err != nil {
		log.Fatalf("failed to register: %v", err)
	}

	err = http.ListenAndServe(proxyAddress, mux)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}