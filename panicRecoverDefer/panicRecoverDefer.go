// panic & recover & defer

package main

import (
	"fmt"
)

func f() (i int) {
	//g()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			i = -1
		}
	}()
	g()
	return 100
}

func g() {
	panic("PANIC!!!")
}

func main() {
	fmt.Println("f() = ", f())
}
