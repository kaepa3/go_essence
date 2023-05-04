package hsd

import (
	"log"
	"os"
	"runtime"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("Before")
	ret := m.Run()
	log.Println("After")
	os.Exit(ret)
}

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
		{name: "lhs is longer than rhs", args: args{lhs: "foo", rhs: "fo"}, want: -1},
		{name: "rhs is shoter than rhs", args: args{lhs: "fo", rhs: "foo"}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringDistance(tt.args.lhs, tt.args.rhs); got != tt.want {
				t.Errorf("StringDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOSing(t *testing.T) {
	log.Println("os test")
	if runtime.GOOS != "windows" {
		t.Skipf("skipping %v", runtime.GOOS)
	}
}

func TestShorter(t *testing.T) {
	log.Println("os test")
	if testing.Short() {
		t.SkipNow()
	}
	log.Println("shoter running")
}
