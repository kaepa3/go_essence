package main

import (
	"fmt"
	"math/rand"
)

func gene(msg string, quit chan string) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			select {
			case ch <- fmt.Sprintf("%s", msg):
			case <-quit:
				quit <- "Seeyou"
				return
			}
		}

	}()
	return ch
}

func main() {
	quit := make(chan string)
	ch := gene("hi", quit)
	for i := rand.Intn(50); i >= 0; i-- {
		fmt.Println(<-ch, i)
	}
	quit <- "Bye"
	fmt.Printf("gene says %s", <-quit)
}
