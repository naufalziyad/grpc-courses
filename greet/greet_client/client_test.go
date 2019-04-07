package main

import (
	"testing"

	"../greetpb"
)

func Test_doServerStreaming(t *testing.T) {
	type args struct {
		c greetpb.GreetServiceClient
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doServerStreaming(tt.args.c)
		})
	}
}

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
