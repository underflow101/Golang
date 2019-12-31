// JSON Unserialization (역직렬화)

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

func unmarshalJSON(b []byte) Task {
	t := Task{}
	err := json.Unmarshal(b, &t)
	if err != nil {
		log.Println(err)
		t.Title = "ERROR"
		t.Status = UNKNOWN
		return t
	}
	return t
}

func main() {
	b := []byte(`{"Title":"Buy Milk","Status":2,"Deadline":"2020-05-26T23:59:59Z"}`)
	t := unmarshalJSON(b)
	fmt.Println(t.Title)
	fmt.Println(t.Status)
	fmt.Println(t.Deadline.UTC())
}
