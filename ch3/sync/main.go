package main

import (
	"fmt"
	"sync"
)

func foring() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		v := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(v)
		}()
	}
	wg.Wait()

}
func ranger() {
	objs := []int{1, 2, 3}
	for i := range objs {
		fmt.Println(i)
	}
}

func racer() {
	n := 0
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			mu.Lock()
			n++
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			mu.Lock()
			n++
			mu.Unlock()
		}
	}()
	wg.Wait()
	fmt.Println(n)
}
func server(ch chan string) {
	defer close(ch)
	ch <- "one"
	ch <- "two"
	ch <- "three"
}

func ch() {
	var s string
	ch := make(chan string)
	go server(ch)
	s = <-ch
	fmt.Println(s)
	s = <-ch
	fmt.Println(s)
	s = <-ch
	fmt.Println(s)
}

func main() {
	foring()
	fmt.Println("----")
	ranger()
	fmt.Println("----")
	racer()
	fmt.Println("----")
	ch()

}
