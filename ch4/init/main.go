//go:generate stringer -type Fruit main.go
package main

import (
	"fmt"
)

var name = "John"

type F struct {
	Name string
	Age  int
}

func (f *F) String() string {
	return fmt.Sprintf("Name=%q, Age=%d", f.Name, f.Age)
}

func main() {
	println("Hi!" + name)
	f := &F{
		Name: "John",
		Age:  20,
	}
	fmt.Printf("%v\n", f)
	fmt.Printf("%+v\n", f)
	fmt.Printf("%#v\n", f)
	fmt.Printf("%T\n", f)

}

func init() {
	println("Hello! " + name)
}

type Fruit int

const (
	Apple Fruit = iota
	Orange
	Banana
)
