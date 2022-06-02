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
	if len(test) == 0 {
		fmt.Println(count)
		return
	}
	for _, word := range test {
		if word == ' ' {
			count++
		}
	}
	count++
	fmt.Println(count)
}
