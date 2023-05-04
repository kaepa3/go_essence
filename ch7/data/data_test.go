package data

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDoSomething(t *testing.T) {
	fns, err := filepath.Glob("testdata/*.dat")
	if err != nil {
		t.Fatal(err)
	}

	for _, fn := range fns {
		t.Log(fn)
		b, err := ioutil.ReadFile(fn)
		if err != nil {
			t.Fatal(err)
		}

		got := doSomething(string(b))

		b, err = ioutil.ReadFile(fn[:len(fn)-3] + "out")
		if err != nil {
			t.Fatal(err)
		}
		want := string(b)
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf(diff)
		}
	}
}

func doSomething(val string) string {
	return val + "out"
}
