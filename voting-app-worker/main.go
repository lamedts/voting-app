package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"voting-app-worker/datastore"
	pb "voting-app-worker/pb"

	"github.com/golang/glog"
	"google.golang.org/grpc"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: example -stderrthreshold=[INFO|WARN|FATAL] -log_dir=[string]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func init() {
	flag.Usage = usage
	flag.Parse()
}

func Run() error {
	// create db connection
	datastore.PgDBInstance = datastore.NewPgDB()
	// create a listener
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		return err
	}
	// create a server instance
	s := pb.EchoServer{}
	grpcServer := grpc.NewServer()
	pb.RegisterEchoServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(listen); err != nil {
		return err
	}
	return nil
}

func main() {
	defer glog.Flush()

	if err := Run(); err != nil {
		glog.Fatal(err)
	}
}
