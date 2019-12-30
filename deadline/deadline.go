package main

import (
	"fmt"
	"time"
)

type Deadline time.Time

func (d Deadline) OverDue() bool {
	return time.Time(d).Before(time.Now())
}

func main() {
	d1 := Deadline(time.Now().Add(-4 * time.Hour))
	d2 := Deadline(time.Now().Add(4 * time.Hour))
	fmt.Println(d1.OverDue())
	fmt.Println(d2.OverDue())
}
