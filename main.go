package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	pb "redModel/pb"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 80, "The server port")
)

type server struct {
	pb.UnimplementedProductServiceServer
}

func main() {
	log.Println("Starting Product Service server...")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProductServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) GetProduct(ctx context.Context, in *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	log.Printf("Received: %v", in.GetId())
	return &pb.GetProductResponse{Product: &pb.Product{}}, nil
}
