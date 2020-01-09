// Go Fib

package main

import "fmt"

func Fib(max uint64) <-chan uint64 {
	ch := make(chan uint64)
	go func() {
		defer close(ch)
		var a, b uint64
		a, b = 0, 1
		for a <= max {
			ch <- a
			a, b = b, a+b
		}
	}()
	return ch
}

func FibGenerator(max uint64) func() uint64 {
	var next, a, b uint64
	next, a, b = 0, 0, 1
	return func() uint64 {
		next, a, b = a, b, a+b
		if next > max {
			return 0
		}
		return next
	}
}

func main() {
	for fib := range Fib(9999999999999999999) {
		fmt.Print(fib, ", ")
	}
	fmt.Println("END!")
}
