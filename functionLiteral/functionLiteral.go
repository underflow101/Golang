package main

import "fmt"

func main() {
	//Annonymous Function
	//Function Literal
	fun := func(a int, b int) int {
		s := a + b
		return s
	}

	fmt.Println(fun(3, 22))
}
