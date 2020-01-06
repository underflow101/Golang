// Custom Printer

package main
import (
	"bufio"
	"fmt"
	"time"
	"encoding/json"
)

type status int
type Deadline struct {
	time.Time
}

func NewDeadline(t time.Time) *Deadline {
	return &Deadline{t}
}

const (
	UNKNOWN = iota
	TODO
	DONE
)

type Task struct {
	Title string `json:"title,omitempty"`
	Status status `json:"status,omitempty"`
	Deadline *Deadline `json:"deadline,omitempty"`
	Priority int `json:"priority,omitempty"`
	SubTasks []Task `json:"subTasks,omitempty"`
}

type IncludeSubTasks Task

func (t IncludeSubTasks) indentedString(prefix string) string {
	str := prefix + Task(t).String()
	for _, st := range t.SubTasks {
		str += "\n" + IncludeSubTasks(st).indentedString(prefix + " ")
	}
	return str
}

func (t IncludeSubTasks) String() string {
	return t.indentedString("")
}

func (t Task) String() string {
	check := "v"
	if t.Status != DONE {
		check = " "
	}
	return fmt.Sprint("[%s] %s %s", check, t.Title, t.Deadline)
}

func PrintStringer(data fmt.Stringer) {
	fmt.Print(data.String())
}

func main() {
	fmt.Println(IncludeSubTasks(Task{
		Title: "Work",
		Status: DONE,
		Deadline: NewDeadline(time.Date(2020, time.January, 15, 23, 59, 59, 0, time.UTC)),
		Priority: 1,
		SubTasks: []Task{{
			Title: "AI Study",
			Status: TODO,
			Deadline: NewDeadline(time.Date(2020, time.January, 15, 23, 59, 59, 0, time.UTC)),
			Prioirty: 2,
			SubTasks: []Task{{
				Title: "Tensorflow-CNN",
				Status: TODO,
				Deadline: NewDeadline(time.Date(2020, ))
			}
			}
		},
	}
	}))
}