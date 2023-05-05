package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var name string
	var max int

	flag.IntVar(&max, "max", 255, "max value")
	flag.StringVar(&name, "name", "", "my name")

	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}
	for _, arg := range flag.Args() {
		fmt.Println(arg)
	}
}
