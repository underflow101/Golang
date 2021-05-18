// skip if no input, get input when there is input

package main

import "fmt"

func FanIn3(in1, in2, in3 <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			select {
			case n := <-in1:
				fmt.Println(n)
			case n := <-in2:
				fmt.Println(n)
			case n := <-in3:
				fmt.Println(n)
			default:
				fmt.Println("Data Not Ready. Skipping...")
			}
		}
	}()
	return out
}

func main() {
	c1, c2, c3 := make(chan int), make(chan int), make(chan int)
	sendInts := func(c chan<- int, begin, end int) {
		defer close(c)
		for i := begin; i < end; i++ {
			c <- i
		}
	}
	go sendInts(c1, 11, 14)
	go sendInts(c2, 21, 23)
	go sendInts(c3, 31, 35)
	for n := range FanIn3(c1, c2, c3) {
		fmt.Print(n, ", ")
	}
	fmt.Println("END!")
}
