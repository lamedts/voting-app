package main

import (
	"fmt"
	"net"
	"voting-app/voting-app-worker/datastore"
	pb "voting-app/voting-app-worker/pb"
	"voting-app/voting-app-worker/utils/logger"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var appLogger *logrus.Entry = logger.GetLogger("app")

func Run() error {
	appLogger.Info("grpc: 50051")
	// create db connection
	datastore.PgDBInstance = datastore.NewPgDB()
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
