package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	pb "redModel/pb"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	// log.Printf("Received: %v %v %v %v", in.GetE().GetId(), in.GetE().GetName(), in.GetE().GetDescription(), in.GetE().GetTime())
	r := connectRedis()
	// log.Println(in.GetE().GetId())
	err := r.HSet(in.GetE().GetId(), "id", in.GetE().GetId()).Err()
	if err != nil {
		log.Println(err)
		return &pb.SetEventResponse{Ok: false}, err
	}
	// log.Println(in.GetE().GetName())
	err = r.HSet(in.GetE().GetId(), "name", in.GetE().GetName()).Err()
	if err != nil {
		log.Println(err)
		return &pb.SetEventResponse{Ok: false}, err
	}
	// log.Println(in.GetE().GetDescription())
	err = r.HSet(in.GetE().GetId(), "description", in.GetE().GetDescription()).Err()
	if err != nil {
		log.Println(err)
		return &pb.SetEventResponse{Ok: false}, err
	}
	// log.Println(in.GetE().GetTime().AsTime())
	// log.Println(in.GetE().GetTime().Seconds)
	err = r.HSet(in.GetE().GetId(), "time", in.GetE().GetTime().Seconds).Err()
	if err != nil {
		log.Println(err)
		return &pb.SetEventResponse{Ok: false}, err
	}
	// log.Println("current time: " + time.Now().String())
	return &pb.SetEventResponse{Ok: true}, nil
}

func (s *server) GetEvent(ctx context.Context, in *pb.GetEventRequest) (*pb.GetEventResponse, error) {
	// log.Printf("Received: %v %v %v %v", in.GetE().GetId(), in.GetE().GetName(), in.GetE().GetDescription(), in.GetE().GetTime())
	event := &pb.Event{}
	r := connectRedis()
	id, err := r.HGet(in.GetE().GetId(), "id").Result()
	if err != nil {
		log.Println(err)
		return &pb.GetEventResponse{E: &pb.Event{}}, err
	}
	event.Id = id
	event.Name, err = r.HGet(in.GetE().GetId(), "name").Result()
	if err != nil {
		log.Println(err)
		return &pb.GetEventResponse{E: &pb.Event{}}, err
	}
	event.Description, err = r.HGet(in.GetE().GetId(), "description").Result()
	if err != nil {
		log.Println(err)
		return &pb.GetEventResponse{E: &pb.Event{}}, err
	}
	time, err := r.HGet(in.GetE().GetId(), "time").Result()
	if err != nil {
		log.Println(err)
		return &pb.GetEventResponse{E: &pb.Event{}}, err
	}
	t, err := stringToTime(time)
	if err != nil {
		log.Println(err)
		return &pb.GetEventResponse{E: &pb.Event{}}, err
	}
	event.Time = &timestamppb.Timestamp{Seconds: t.Unix()}
	return &pb.GetEventResponse{E: event}, nil
}

// https://stackoverflow.com/questions/43431046/convert-unix-epoch-as-a-string-to-time-time-on-go
func stringToTime(s string) (time.Time, error) {
	sec, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(sec, 0), nil
}

func (s *server) GetEvents(ctx context.Context, in *pb.GetEventsRequest) (*pb.GetEventsResponse, error) {
	r := connectRedis()
	events := []*pb.Event{}
	keys, err := r.Keys("*").Result()
	if err != nil {
		log.Println(err)
		return &pb.GetEventsResponse{E: events}, err
	}
	for _, key := range keys {
		event := &pb.Event{}
		event.Id = key
		event.Name, err = r.HGet(key, "name").Result()
		if err != nil {
			log.Println(err)
			return &pb.GetEventsResponse{E: events}, err
		}
		event.Description, err = r.HGet(key, "description").Result()
		if err != nil {
			log.Println(err)
			return &pb.GetEventsResponse{E: events}, err
		}
		time, err := r.HGet(key, "time").Result()
		if err != nil {
			log.Println(err)
			return &pb.GetEventsResponse{E: events}, err
		}
		t, err := stringToTime(time)
		if err != nil {
			log.Println(err)
			return &pb.GetEventsResponse{E: events}, err
		}
		event.Time = &timestamp.Timestamp{Seconds: t.Unix()}
		events = append(events, event)
	}
	return &pb.GetEventsResponse{E: events}, nil
}
