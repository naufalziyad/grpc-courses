package main

import (
	"testing"

	"../greetpb"
)

func Test_doUnary(t *testing.T) {
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
			doUnary(tt.args.c)
		})
	}
}
