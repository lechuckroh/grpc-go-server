package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/lechuckroh/grpc-go-server/config"
	"github.com/lechuckroh/grpc-go-server/handler"
	"github.com/lechuckroh/grpc-go-server/pb/hellopb"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

// runGateway starts gRPC-gateway
func runGateway(cfg *config.APIConfig) {
	httpPort := cfg.HTTP.Port
	if httpPort <= 0 {
		log.Print("gRPC-Gateway is disabled")
		return
	}

	ctx := context.Background()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	grpcServerEndpoint := fmt.Sprintf("localhost:%d", cfg.GRPC.Port)
	if err := hellopb.RegisterHelloGwFromEndpoint(ctx, mux, grpcServerEndpoint, opts); err != nil {
		log.Fatal("failed to register hello gateway handlers", err)
	}

	log.Printf("starting HTTP gateway: TCP %s", grpcServerEndpoint)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", httpPort), mux); err != nil {
		log.Fatal("failed to start gateway", err)
	}
}

func main() {
	// Load configuration
	cfg := config.LoadConfig("")

	address := fmt.Sprintf(":%d", cfg.GRPC.Port)
	tcpListener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	hellopb.RegisterHelloServer(grpcServer, &handler.Hello{})

	go runGateway(cfg)

	log.Printf("starting gRPC server: TCP %s", address)
	if err := grpcServer.Serve(tcpListener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
