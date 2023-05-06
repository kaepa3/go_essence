package hsd

import "testing"

func TestStringDistance(t *testing.T) {
	type args struct {
		lhs string
		rhs string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "basictest", args: args{lhs: "foo", rhs: "foh"}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringDistance(tt.args.lhs, tt.args.rhs); got != tt.want {
				t.Errorf("StringDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
