package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/oskoss/t1/api"
	"google.golang.org/grpc"
)

func startGRPCServer(address string) error {
	// create a listener on TCP port
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	// create a DeploymentService instance
	s := api.DeploymentService{}
	grpcServer := grpc.NewServer()
	api.RegisterDeploymentServer(grpcServer, &s)
	log.Printf("starting HTTP/2 gRPC server on %s", address)
	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %s", err)
	}
	return nil
}

func startRESTServer(address, grpcAddress string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	// Setup the client gRPC options
	// Register ping
	err := api.RegisterDeploymentHandlerFromEndpoint(ctx, mux, grpcAddress, []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		return fmt.Errorf("could not register DeploymentService: %s", err)
	}
	log.Printf("starting HTTP/1.1 REST server on %s", address)
	http.ListenAndServe(address, mux)
	return nil
}

func main() {

	grpcAddress := fmt.Sprintf("%s:%d", "localhost", 7777)
	restAddress := fmt.Sprintf("%s:%d", "localhost", 7778)
	// fire the gRPC server in a goroutine
	go func() {
		err := startGRPCServer(grpcAddress)
		if err != nil {
			log.Fatalf("failed to start gRPC server: %s", err)
		}
	}()
	// fire the REST server in a goroutine
	go func() {
		err := startRESTServer(restAddress, grpcAddress)
		if err != nil {
			log.Fatalf("failed to start REST server: %s", err)
		}
	}()
	// infinite loop
	log.Printf("Entering infinite loop")
	select {}
}
