package main

import (
	"fmt"
	"net"
	"voting-app/voting-app-worker/config"
	"voting-app/voting-app-worker/datastore"
	pb "voting-app/voting-app-worker/pb"
	"voting-app/voting-app-worker/utils/logger"

	"google.golang.org/grpc"
)

var appLogger = logger.GetLogger("app")

func Run() error {
	// create db connection
	datastore.PgDBInstance = datastore.NewPgDB(config.AppConfig.PgDB)
	// create a listener
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		return err
	}
	appLogger.Info("grpc: 50051")
	// create a server instance
	s := pb.VoteServer{}
	grpcServer := grpc.NewServer()
	pb.RegisterVoteWorkerServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(listen); err != nil {
		return err
	}
	return nil
}

func main() {

	appLogger.Info("start")
	if err := Run(); err != nil {
		appLogger.Fatal(err)
	}
}
