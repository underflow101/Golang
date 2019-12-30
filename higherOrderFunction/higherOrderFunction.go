package main

import "fmt"

func main() {
	tmpFunc := func(a int, b int, c int) int {
		s := a + b*c
		return s
	}
	tmpFunc2 := func(a int, b int) int {
		s := a + b
		return s
	}
	res := tmpFunc(tmpFunc2(1, 2), tmpFunc2(5, 6), tmpFunc2(9, 10))
	fmt.Println(res)
}
