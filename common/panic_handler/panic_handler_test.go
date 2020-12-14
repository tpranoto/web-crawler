package panics

import (
	"log"
	"testing"
)

func TestHandlePanic(t *testing.T) {
	type args struct {
		loc string
	}
	tests := []struct {
		name  string
		args  args
		setup func()
	}{
		{
			name: "valid",
			args: args{
				loc: "[Test]",
			},
		},
		{
			name: "panic!",
			args: args{
				loc: "[Test]",
			},
			setup: func() {
				panic("test")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer HandlePanic(tt.args.loc)
			if tt.setup != nil {
				tt.setup()
			}
		})
	}
}

func TestConcurrentHandlePanic(t *testing.T) {
	type args struct {
		loc       string
		handlerFn func()
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "valid",
			args: args{
				loc: "[Test]",
				handlerFn: func() {
					log.Println("test")
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ConcurrentHandlePanic(tt.args.loc, tt.args.handlerFn)
		})
	}
}
