package calc

import "testing"

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{a: 0, b: 1}, want: 1},
		{name: "test2", args: args{a: 1, b: -1}, want: 0},
		{name: "test3", args: args{a: 2, b: 1}, want: 3},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
