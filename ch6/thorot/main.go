package main

import (
	"fmt"
	"sync"
	"time"
)

func downloadJSON(u string) {
	println(u)
	time.Sleep(time.Second * 6)
}

func main() {
	before := time.Now()

	limit := make(chan struct{}, 20)

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		i := i
		go func() {
			limit <- struct{}{}

			defer wg.Done()
			u := fmt.Sprintf("http://example.com/api/user?id=%d", i)
			downloadJSON(u)
			<-limit
		}()
	}
	wg.Wait()
	fmt.Printf("%v\n", time.Since(before))
}
