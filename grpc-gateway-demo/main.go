package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	server "github.com/soypete/Golang_datatools_exercises/grpc-gateway-demo/server"
)

var (
	httpPort string
	grpcPort string
)

func main() {
	flag.StringVar(&httpPort, "httpPort", ":8090", "HTTP Port to listen on")
	flag.StringVar(&grpcPort, "grpcPort", ":9090", "gRPC Port to listen on")
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	grpcServer := server.SetupGrpc()

	fmt.Println("gRPCServer is configured and listening on port :9090")

	go grpcServer.RunGrpc(ctx, grpcPort)

	// Configure Gateway Server
	grpcServer.SetupGateway(ctx, httpPort, grpcPort)
	fmt.Println("GatewayServer is configured and running on port :8090")
	err := grpcServer.GWServer.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
	// system close call to close all connections
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		sig := <-c
		fmt.Printf("Got %s signal. Aborting...\n", sig)
		grpcServer.Close(ctx)
		os.Exit(1)
	}()

}
