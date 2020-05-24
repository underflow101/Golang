// Channel to communicate within go routines
// Can be used in queue mechanism

package main

import "fmt"

func main() {
	// channel must me made by make() function before hand
	// use <- operator to send/receive data
	ch := make(chan int)
	done := make(chan bool)

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		ch <- 4
		ch <- 5
	}()

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(<-ch)
		}
		done <- true
	}()

	<-done
	fmt.Println("Done!")
}
