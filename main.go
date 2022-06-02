package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		return
	}
	test := os.Args[1]
	var count int
	for _, word := range test {
		if word == ' ' {
			count++
		}
	}
	count++
	fmt.Println(count)
}
