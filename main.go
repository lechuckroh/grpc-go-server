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
	"os"
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

func runServer(cfg *config.APIConfig) {
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

func runClient(cfg *config.APIConfig) {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
	}

	// Connect
	serverAddr := fmt.Sprintf("127.0.0.1:%d", cfg.GRPC.Port)
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("failed to connect server: %s", serverAddr)
	}
	defer func() {
		_ = conn.Close()
	}()

	// create client
	client := hellopb.NewHelloClient(conn)
	ctx := context.Background()

	// call
	res, err := client.Call(ctx, &hellopb.CallRequest{Name: "World"})
	if err != nil {
		log.Fatalf("failed to run CallRequest: %v", err)
	}

	msg := res.GetMsg()
	log.Printf("CallRequest reponse: %s", msg)
}

func main() {
	// Load configuration
	cfg := config.LoadConfig("")

	runMode := os.Getenv("RUN_MODE")
	if runMode == "client" {
		runClient(cfg)
	} else {
		runServer(cfg)
	}
}
