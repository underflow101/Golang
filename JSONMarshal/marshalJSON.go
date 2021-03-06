// JSON Serialization (직렬화)

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type status int
type Deadline struct {
	time.Time
}

const (
	UNKNOWN = iota
	TODO
	DONE
)

func NewDeadline(t time.Time) *Deadline {
	return &Deadline{t}
}

type Task struct {
	Title    string
	Status   status
	Deadline *Deadline
}

func marshalJSON(t Task) string {
	b, err := json.Marshal(t)
	if err != nil {
		log.Println(err)
		return "Error"
	}
	return string(b)
}

func main() {
	t := Task{
		"Laundry",
		DONE,
		NewDeadline(time.Date(2020, time.January, 16, 23, 59, 59, 0, time.UTC)),
	}
	b := marshalJSON(t)
	fmt.Println(b)
}
