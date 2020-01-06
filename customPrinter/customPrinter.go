// Custom Printer

package main
import (
	"fmt"
	"time"
	//"encoding/json"
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
	if t.Status == TODO {
		check = " "
	} else if t.Status == UNKNOWN {
		check = "?"
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
			Priority: 3,
			SubTasks: []Task{{
				Title: "Tensorflow-CNN",
				Status: DONE,
				Deadline: NewDeadline(time.Date(2020, time.January, 13, 23, 59, 59, 0, time.UTC)),
				Priority: 1,
			}, {
				Title: "Keras",
				Status: TODO,
				Deadline: NewDeadline(time.Date(2020, time.January, 14, 23, 59, 59, 0, time.UTC)),
				Priority: 2,
			}}},
			{
				Title: "Golang Study",
				Status: UNKNOWN,
				Deadline: NewDeadline(time.Date(2020, time.January, 8, 23, 59, 59, 0, time.UTC)),
				Priority: 1,
			}, {
				Title: "MQTT",
				Status: DONE,
				Deadline: NewDeadline(time.Date(2020, time.January, 10, 23, 59, 59, 0, time.UTC)),
				Priority: 2,
			}}}))
	// fmt.Println(IncludeSubTasks(Task{
	// 	Title: "Life",
	// 	Status: UNKNOWN,
	// 	Deadline: NewDeadline(time.Date(2150, time.January, 31, 23, 59, 59, 0, time.UTC())),
	// 	Priority: 999999999999,
	// 	SubTasks: []Task{
	// 		Title: "is Death",
	// 		Status: DONE,
	// 		Deadline: NewDeadline(time.Date(2150, time.January, 31, 23, 59, 59, 0, time.UTC)),
	// 		Priority: 1,
	// }}))
}