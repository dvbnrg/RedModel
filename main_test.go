package main

import (
	"context"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	pb "redModel/pb"
	"reflect"
	"testing"
)

func Test_server_GetProduct(t *testing.T) {
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
		{"test1", fields{}, args{context.Background(), &pb.GetEventRequest{E: &pb.Event{Id: "1", Name: "testEvent", Description: "test Event Description", Time: &timestamppb.Timestamp{}}}}, &pb.GetEventResponse{E: &pb.Event{}}, false},
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
