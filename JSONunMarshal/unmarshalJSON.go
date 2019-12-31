// JSON Unserialization (역직렬화)

package main
import (
	"fmt"
	"JSON"
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

func main() {
	
}