// Pipeline Pattern

package main

import "fmt"

func PlusOne(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			out <- (num + 1)
		}
	}()
	return out
}

func main() {
	c := make(chan int)
	go func() {
		defer close(c)
		c <- 5
		c <- 3
		c <- 8
	}()
	for num := range PlusOne(PlusOne(PlusOne(c))) {
		fmt.Println(num)
	}
}
