// race condition solver

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var cnt int = 10
	var n int = 1
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			cnt--
			switch n {
			case 1:
				fmt.Println("I'm", n, "st Go Routine! :D and now cnt is:", cnt)
			case 2:
				fmt.Println("I'm", n, "nd Go Routine! :D and now cnt is:", cnt)
			case 3:
				fmt.Println("I'm", n, "rd Go Routine! :D and now cnt is:", cnt)
			default:
				fmt.Println("I'm", n, "th Go Routine! :D and now cnt is:", cnt)
			}
		}(n)
		n++
	}
	wg.Wait()
	fmt.Println("cnt is now: ", cnt)
}
