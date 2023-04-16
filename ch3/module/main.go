package main

import (
	"fmt"

	"github.com/mattn/go-runewidth"
)

func main() {
	fmt.Println(runewidth.StringWidth("こんにちは"))
}
