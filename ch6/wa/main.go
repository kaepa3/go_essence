package main

import (
	"fmt"
	"time"
)

func gene(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Second)
		}
	}()
	return ch
}

func main() {
	ch := gene("hi")
	to := time.After(5 * time.Second)
	for i := 0; i < 10; i++ {
		select {
		case s := <-ch:
			fmt.Println(s)
		case <-to:
			fmt.Println("5s timeout")
			return
		}
	}
}
