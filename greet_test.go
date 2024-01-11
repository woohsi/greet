package greet

import (
	"testing"
)

func TestHello(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test me",
			args{"xiang"},
			"Hi [xiang]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(tt.args.name); got != tt.want {
				t.Errorf("Hello(%v) = %v, want %v", tt.args.name, got, tt.want)
			}
		})
	}
}
