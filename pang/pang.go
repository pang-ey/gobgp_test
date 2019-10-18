package pang

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	pb "gobgp/onos"

	"google.golang.org/grpc"
)

const (
	address     = "192.168.232.142:50052"
	defaultName = "Nothing happens"
)

// send message by grpc
func Helloworld(param string) {
	// fmt.Println("hello pang")
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		// log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewOnosServerClient(conn)

	name := defaultName
	if len(param) > 0 {
		name = param
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.EvpnRoute(ctx, &pb.OnosRequest{Message: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Info("Reply: ", r.Message)
}
