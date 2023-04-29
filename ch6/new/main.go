package main

import (
	"fmt"
	"time"
)

func fanin(ch1, ch2 <-chan string) <-chan string {
	new_ch := make(chan string)
	go func() {
		for {
			select {
			case s := <-ch1:
				new_ch <- s
			case s := <-ch2:
				time.Sleep(time.Second)
				new_ch <- s
			}
		}
	}()
	return new_ch
}

func generator(msg string) <-chan string {

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
	ch := fanin(generator("Hello"), generator("bye"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}
