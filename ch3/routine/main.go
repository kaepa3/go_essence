package main

import "fmt"

func main() {
	messege := "hi"
	go sendMessage(messege)
}

func sendMessage(msg string) {
	fmt.Println(msg)
}
