// Dynamic Go Routine Connection (Prime Number)

package main

import (
	"context"
	"fmt"
	"runtime"
)

type IntPipe func(context.Context, <-chan int) <-chan int

var res []int

func Range(ctx context.Context, start, step int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := start; ; i += step {
			select {
			case out <- i:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

func FilterMultiple(n int) IntPipe {
	return func(ctx context.Context, in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for x := range in {
				if x%n == 0 {
					continue
				}
				select {
				case out <- x:
				case <-ctx.Done():
					return
				}
			}
		}()
		return out
	}
}

func Primes(ctx context.Context) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		c := Range(ctx, 2, 1)
		for {
			select {
			case i := <-c:
				c = FilterMultiple(i)(ctx, c)
				select {
				case out <- i:
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

func PrintPrimes(max int) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for prime := range Primes(ctx) {
		if prime > max {
			break
		}
		res = append(res, prime)
		//fmt.Print(prime, " ")
	}
	fmt.Println()
}

func main() {
	x := runtime.NumCPU()
	fmt.Println(x)
	runtime.GOMAXPROCS(x)
	PrintPrimes(1000)
	fmt.Println(res)
}
