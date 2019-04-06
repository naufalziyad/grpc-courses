package main

import (
	"context"
	"reflect"
	"testing"

	"../greetpb"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_server_Greet(t *testing.T) {
	type args struct {
		ctx context.Context
		req *greetpb.GreetRequest
	}
	tests := []struct {
		name    string
		s       *server
		args    args
		want    *greetpb.GreetResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{}
			got, err := s.Greet(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.Greet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.Greet() = %v, want %v", got, tt.want)
			}
		})
	}
}
