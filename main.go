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

func connectRedis() (r *redis.Client) {
	godotenv.Load(".env")
	url := os.Getenv("REDIS")
	client := redis.NewClient(&redis.Options{Addr: url, Password: "", DB: 0})
	// pong, err := client.Ping().Result()
	// log.Println(pong, err)
	return client
}

func (s *server) SetEvent(ctx context.Context, in *pb.SetEventRequest) (*pb.SetEventResponse, error) {
	log.Printf("Received: %v %v %v %v", in.GetE().GetId(), in.GetE().GetName(), in.GetE().GetDescription(), in.GetE().GetTime())
	r := connectRedis()
	log.Println(in.GetE().GetId())
	err := r.HSet(in.GetE().GetId(), "id", in.GetE().GetId()).Err()
	if err != nil {
		log.Println(err)
		return &pb.SetEventResponse{Ok: false}, err
	}
	log.Println(in.GetE().GetName())
	err = r.HSet(in.GetE().GetId(), "name", in.GetE().GetName()).Err()
	if err != nil {
		log.Println(err)
		return &pb.SetEventResponse{Ok: false}, err
	}
	log.Println(in.GetE().GetDescription())
	err = r.HSet(in.GetE().GetId(), "description", in.GetE().GetDescription()).Err()
	if err != nil {
		log.Println(err)
		return &pb.SetEventResponse{Ok: false}, err
	}
	log.Println(in.GetE().GetTime().AsTime())
	log.Println(in.GetE().GetTime().Seconds)
	err = r.HSet(in.GetE().GetId(), "time", in.GetE().GetTime().AsTime()).Err()
	if err != nil {
		log.Println(err)
		return &pb.SetEventResponse{Ok: false}, err
	}
	log.Println("current time: " + time.Now().String())
	return &pb.SetEventResponse{Ok: true}, nil
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
