package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	pb "helloapp/proto/helloworld"
)

var gRPCServerAddr = flag.String("grpc.server.addr", "helloapp_nginx_lb:50052", "grpc server addr")

func main() {
	conn, err := grpc.Dial(*gRPCServerAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	name := "world"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not hello: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
