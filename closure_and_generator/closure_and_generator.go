package main

//Closure of Golang
func Closure_nextVal() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {
	//Generator
	next := Closure_nextVal()
	println(next())
	println(next())
	println(next())

	wowNext := Closure_nextVal()
	println(wowNext())
	println(wowNext())
}
