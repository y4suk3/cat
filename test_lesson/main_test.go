package main

import (
	"fmt"
	"testing"
)

func Test_add(t *testing.T) {
	t.Parallel()
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{x: 1, y: 3}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_greeting(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test2", args: args{s: "gopher"}, want: "Hi, gopher"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := greeting(tt.args.s); got != tt.want {
				t.Errorf("greeting() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Example_greeting() {
	fmt.Println(greeting("kore"))

	// Output:
	// Hi, kore
}
