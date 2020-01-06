package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 20; i++ {
		fmt.Println(i)
	}
	s := "Hello, World!"
	fmt.Printf("[%s]", s)
}
