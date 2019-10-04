package main

import (
	"context"
	"log"
	"time"

	pb "github.com/pang-ey/gobgp_test/api/OnosServer"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50052"
	defaultName = "Nothing happens"
)

func Helloworld() {
	// fmt.Println("hello pang")
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		// log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewOnosServerClient(conn)

	name := defaultName

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.EvpnRoute(ctx, &pb.OnosRequest{Message: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Reply: %s", r.Message)
}
