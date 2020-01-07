// go routine basic

package main

import "fmt"

func main() {
	go func() {
		fmt.Println("I'm in Go Routine!!! :)")
	}()
	fmt.Println("I'm in Main Routine, Damn it!!! :(")
}
