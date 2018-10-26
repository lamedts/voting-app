package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"voting-app/voting-app-worker/config"
	"voting-app/voting-app-worker/datastore"
	"voting-app/voting-app-worker/pb"
	"voting-app/voting-app-worker/utils/logger"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

var appLogger = logger.GetLogger("app")

func Run(grpcAddress string, restAddress string) {

	go func(grpcAddress string) {
		lis, err := net.Listen("tcp", grpcAddress)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		s := pb.VoteServer{}
		grpcServer := grpc.NewServer()
		pb.RegisterVoteWorkerServiceServer(grpcServer, &s)

		appLogger.Printf("starting HTTP/2 gRPC server on %s", grpcAddress)
		if err := grpcServer.Serve(lis); err != nil {
			appLogger.Fatalf("failed to serve: %s", err)
		}
	}(grpcAddress)

	go func(grpcAddress string, restAddress string) {
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		mux := runtime.NewServeMux()
		opts := []grpc.DialOption{grpc.WithInsecure()}
		// Register ping
		err := pb.RegisterVoteWorkerServiceHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
		if err != nil {
			appLogger.Fatalf("could not register service: %s", err)
		}

		appLogger.Printf("starting HTTP/1.1 REST server on %s", restAddress)
		http.ListenAndServe(restAddress, mux)
	}(grpcAddress, restAddress)

}

func main() {
	appLogger.Info("start")

	// create db connection
	datastore.PgDBInstance = datastore.NewPgDB(config.AppConfig.PgDB)

	restPoert := 50052
	gRPCPoert := 50051
	Run(fmt.Sprintf(":%d", gRPCPoert), fmt.Sprintf(":%d", restPoert))
	select {}
}
