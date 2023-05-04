package fuzz

import (
	"log"
	"testing"
)

func FuzzDoSomthing(f *testing.F) {
	f.Add("test&&&")
	f.Fuzz(func(f *testing.T, s string) {
		doSomething(s)
	})
}

func doSomething(val string) {
	log.Println(val)
}
