// Closure of Golang

package main

func Closure_nextVal() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {
	// closure
	next := Closure_nextVal()
	println(next())
	println(next())
	println(next())
	
	// closure
	wowNext := Closure_nextVal()
	println(wowNext())
	println(wowNext())
}
