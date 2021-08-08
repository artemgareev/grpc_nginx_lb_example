package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	pb "helloapp/proto/helloworld"
)

var gRPCPort = flag.String("grpc.port", ":50051", "grpc port")

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	hostName, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	return &pb.HelloReply{Message: "Hello " + in.GetName() + " " + hostName}, nil
}

func main() {
	lis, err := net.Listen("tcp", *gRPCPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
