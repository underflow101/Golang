// making queue with Go channel

package main

import "fmt"

type Queue struct {
	item chan int
}

// save inputs
func (q *Queue) Append(val int) {
	q.item <- val
}

// pop inputs
func (q *Queue) Pop() int {
	return <-q.item
}

func main() {
	q := Queue{item: make(chan int, 5)}
	q.Append(1)
	q.Append(2)
	q.Append(3)

	a := q.Pop()
	b := q.Pop()
	c := q.Pop()

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
