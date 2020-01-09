// Pipeline Pattern

package main

import "fmt"

type IntPipe func(<-chan int) <-chan int

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

func Chain(ps ...IntPipe) IntPipe {
	return func(in <-chan int) <-chan int {
		c := in
		for _, p := range ps {
			c = p(c)
		}
		return c
	}
}

func main() {
	c := make(chan int)
	//PlusTwo := Chain(PlusOne, PlusOne, PlusOne)
	PlusFive := Chain(PlusOne, PlusOne, PlusOne, PlusOne, PlusOne)
	go func() {
		defer close(c)
		c <- 5
		c <- 3
		c <- 8
	}()
	// for num := range PlusTwo(c) {
	// 	fmt.Println(num)
	// }
	for n := range PlusFive(c) {
		fmt.Println(n)
	}
}
