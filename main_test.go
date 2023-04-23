package main

import (
	"context"
	pb "redModel/pb"
	"reflect"
	"testing"
)

func Test_server_GetProduct(t *testing.T) {
	type fields struct {
		UnimplementedProductServiceServer pb.UnimplementedProductServiceServer
	}
	type args struct {
		ctx context.Context
		in  *pb.GetProductRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.GetProductResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{"test1", fields{}, args{context.Background(), &pb.GetProductRequest{Id: "1"}}, &pb.GetProductResponse{Product: &pb.Product{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				UnimplementedProductServiceServer: tt.fields.UnimplementedProductServiceServer,
			}
			got, err := s.GetProduct(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.GetProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.GetProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
