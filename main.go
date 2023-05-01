package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	pb "redModel/pb"
	"time"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 80, "The server port")
)

type server struct {
	pb.UnimplementedEventServiceServer
}

func main() {
	log.Println("Starting Event Service server...")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEventServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) SetEvent(ctx context.Context, in *pb.SetEventRequest) (*pb.SetEventResponse, error) {
	log.Printf("Received: %v %v %v %v", in.GetE().GetId(), in.GetE().GetName(), in.GetE().GetDescription(), in.GetE().GetTime())
	godotenv.Load(".env")
	r := os.Getenv("REDIS")
	fmt.Println("Testing Golang Redis")
	client := redis.NewClient(&redis.Options{
		Addr:     r,
		Password: "",
		DB:       0,
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	log.Println(time.Now())
	return &pb.SetEventResponse{}, nil
}

func (s *server) GetEvent(ctx context.Context, in *pb.GetEventRequest) (*pb.GetEventResponse, error) {
	log.Printf("Received: %v %v %v %v", in.GetE().GetId(), in.GetE().GetName(), in.GetE().GetDescription(), in.GetE().GetTime())
	log.Println(time.Now())
	return &pb.GetEventResponse{E: &pb.Event{}}, nil
}

func (s *server) GetEvents(ctx context.Context, in *pb.GetEventsRequest) (*pb.GetEventsResponse, error) {
	log.Println(time.Now())
	return &pb.GetEventsResponse{E: []*pb.Event{}}, nil
}
