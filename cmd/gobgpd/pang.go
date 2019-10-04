package main

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	pb "github.com/pang-ey/gobgp_test/onos"

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
	log.Info("Reply: ", r.Message)
}
