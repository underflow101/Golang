//Strings

package main

import "fmt"

func main() {
	//Raw String Literal
	var s string = `This
is
Sparta!!!\n`
	fmt.Println(s)

	//Interpreted String Literal
	var d string = "This\nis\nSparta!!!\nAND\nSomething\nMore?!"
	fmt.Println(d)
}
