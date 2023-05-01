package main

import (
	"context"
	pb "redModel/pb"
	"reflect"
	"testing"

	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Test_server_SetEvent(t *testing.T) {
	type fields struct {
		UnimplementedEventServiceServer pb.UnimplementedEventServiceServer
	}
	type args struct {
		ctx context.Context
		in  *pb.SetEventRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.SetEventResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Set Test Event 1", fields{}, args{context.Background(), &pb.SetEventRequest{E: &pb.Event{Id: "test1", Name: "testEvent1", Description: "test Event 1 Description", Time: &timestamppb.Timestamp{Seconds: 1525138858}}}}, &pb.SetEventResponse{Ok: true}, false},
		{"Set Test Event 2", fields{}, args{context.Background(), &pb.SetEventRequest{E: &pb.Event{Id: "test2", Name: "testEvent2", Description: "test Event 2 Description", Time: &timestamppb.Timestamp{Seconds: 1525138858}}}}, &pb.SetEventResponse{Ok: true}, false},
		{"Set Test Event 3", fields{}, args{context.Background(), &pb.SetEventRequest{E: &pb.Event{Id: "test3", Name: "testEvent3", Description: "test Event 3 Description", Time: &timestamppb.Timestamp{Seconds: 1525138858}}}}, &pb.SetEventResponse{Ok: true}, false},
		{"Set Test Event 4", fields{}, args{context.Background(), &pb.SetEventRequest{E: &pb.Event{Id: "test4", Name: "testEvent4", Description: "test Event 4 Description", Time: &timestamppb.Timestamp{Seconds: 1525138858}}}}, &pb.SetEventResponse{Ok: true}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				UnimplementedEventServiceServer: tt.fields.UnimplementedEventServiceServer,
			}
			got, err := s.SetEvent(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.SetEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.SetEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_GetEvent(t *testing.T) {
	type fields struct {
		UnimplementedEventServiceServer pb.UnimplementedEventServiceServer
	}
	type args struct {
		ctx context.Context
		in  *pb.GetEventRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.GetEventResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Get Test Event 1", fields{}, args{context.Background(), &pb.GetEventRequest{E: &pb.Event{Id: "test1"}}}, &pb.GetEventResponse{E: &pb.Event{Id: "test1", Name: "testEvent1", Description: "test Event 1 Description", Time: &timestamppb.Timestamp{Seconds: 1525138858}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				UnimplementedEventServiceServer: tt.fields.UnimplementedEventServiceServer,
			}
			got, err := s.GetEvent(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.GetEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.GetEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_DeleteEvent(t *testing.T) {
	type fields struct {
		UnimplementedEventServiceServer pb.UnimplementedEventServiceServer
	}
	type args struct {
		ctx context.Context
		in  *pb.DeleteEventRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.DeleteEventResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Delete Test Event 1", fields{}, args{context.Background(), &pb.DeleteEventRequest{E: &pb.Event{Id: "test1"}}}, &pb.DeleteEventResponse{Ok: true}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				UnimplementedEventServiceServer: tt.fields.UnimplementedEventServiceServer,
			}
			got, err := s.DeleteEvent(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.DeleteEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.DeleteEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_GetEvents(t *testing.T) {
	type fields struct {
		UnimplementedEventServiceServer pb.UnimplementedEventServiceServer
	}
	type args struct {
		ctx context.Context
		in  *emptypb.Empty
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.GetEventsResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Get All Test Events", fields{}, args{context.Background(), &emptypb.Empty{}}, &pb.GetEventsResponse{E: []*pb.Event{{Id: "test3", Name: "testEvent3", Description: "test Event 3 Description", Time: &timestamppb.Timestamp{Seconds: 1525138858}}, {Id: "test4", Name: "testEvent4", Description: "test Event 4 Description", Time: &timestamppb.Timestamp{Seconds: 1525138858}}, {Id: "test2", Name: "testEvent2", Description: "test Event 2 Description", Time: &timestamppb.Timestamp{Seconds: 1525138858}}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				UnimplementedEventServiceServer: tt.fields.UnimplementedEventServiceServer,
			}
			got, err := s.GetEvents(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.GetEvents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.GetEvents() = %v, want %v", got, tt.want)
			}
		})
	}
}
