package main

import (
	"io"
	"log"
	pb "voting-app/voting-app-worker/pb"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := pb.NewVoteWorkerServiceClient(conn)
	status, setErr := c.SetVote(context.Background(), &pb.Vote{Vote: "dog", VotedID: 87})
	if setErr != nil {
		log.Fatalf("cannot insert")
	}
	log.Printf("status: %+v", status)

	stream, err := c.GetVotesResults(context.Background(), &pb.WorkerRequest{Query: "foo"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	// log.Printf("Response from server: %s", response)
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			log.Println("and now your watch is ended")
			return
		}
		if err != nil {
			log.Fatal(err)
		}
		log.Println(res)
	}
}
