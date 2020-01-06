// Data Access Object for RESTful API
// Low level version: NOT thread safe and not for use (unable to handle when multiple request comes)
// Only for studying

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

type ID int64
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
	Title    string    `json:"title,omitempty"`
	Status   status    `json:"status,omitempty"`
	Deadline *Deadline `json:"deadline,omitempty"`
	Priority int       `json:"priority,omitempty"`
	SubTasks []Task    `json:"subTasks,omitempty"`
}

type IncludeSubTasks Task

func (t IncludeSubTasks) indentedString(prefix string) string {
	str := prefix + Task(t).String()
	for _, st := range t.SubTasks {
		str += "\n" + IncludeSubTasks(st).indentedString(prefix+" ")
	}
	return str
}

func (t IncludeSubTasks) String() string {
	return t.indentedString("")
}

func (t Task) String() string {
	check := "v"
	if t.Status == TODO {
		check = "-"
	} else if t.Status == UNKNOWN {
		check = "?"
	}
	return fmt.Sprintf("[%s] %s %s", check, t.Title, t.Deadline)
}

func PrintStringer(data fmt.Stringer) {
	fmt.Print(data.String())
}

// Data Access Object
type DataAccess interface {
	Get(id ID) (task.Task, error)
	Put(id ID, t task.Task) error
	Post(t task.Task) (ID, error)
	Delete(id ID) error
}

type MemoryDataAccess struct {
	tasks  map[ID]task.Task
	nextID ID
}

func NewMemoryDataAccess() DataAccess {
	return &MemoryDataAccess{
		tasks:  map[ID]task.Task{},
		nextID: ID(1),
	}
}

var ErrTaskNotExist = errors.New("task does not exist")

func (m *MemoryDataAccess) Get(id ID) (task.Task, error) {
	t, exists := m.tasks[id]
	if !exists {
		return task.Task{}, ErrTaskNotExist
	}
	return t, nil
}

func (m *MemoryDataAccess) Put(id ID, t task.Task) error {
	if _, exists := m.tasks[id]; !exists {
		return ErrTaskNotExist
	}
	m.tasks[id] = t
	return nil
}

func (m *MemoryDataAccess) Post(t task.Task) (ID, error) {
	id := ID(fmt.Sprint(m.nextID))
	m.nextID++
	m.tasks[id] = t
	return id, nil
}

func (m *MemoryDataAccess) Delete(id ID) error {
	if _, exists := m.tasks[id]; !exists {
		return ErrTaskNotExist
	}
	delete(m.tasks, id)
	return nil
}

// RESTful API Handler
type ResponseError struct {
	Err error
}

func (err ResponseError) MarshalJSON() ([]byte, error) {
	if err.Err == nil {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%v\"", err.Err)), nil
}

func (err *ResponseError) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, v); err != nil {
		return err
	}
	if v == nil {
		err.Err = nil
		return nil
	}
	switch tv := v.(type) {
	case string:
		if tv == ErrTaskNotExist.Error() {
			err.Err = ErrTaskNotExist
			return nil
		}
		err.Err = errors.New(tv)
		return nil
	default:
		return errors.New("ResponseError unmarshal failed")
	}
}

// Response struct
type Response struct {
	ID    ID            `json:"id,omitempty"`
	Task  task.Task     `json:"task"`
	Error ResponseError `json:"error"`
}

var m = NewMemoryDataAccess()

const pathPrefix = "/api/v1/task/"

func apiHandler(w http.ResponseWriter, r *http.Request) {
	getID := func() (ID, error) {
		id := task.ID(r.URL.Path[len(pathPrefix):])
		if id == "" {
			return id, errors.New("apiHandler: ID is empty")
		}
		return id, nil
	}

	getTasks := func() ([]task.Task, error) {
		var result []task.Task
		if err := r.ParseForm(); err != nil {
			return nil, err
		}
		encodedTasks, ok := r.PostForm["task"]
		if !ok {
			return nil, errors.New("task parameter expected")
		}
		for _, encodedTask := range encodedTasks {
			var t task.Task
			if err := json.Unmarshal([]byte(encodedTask), &t); err != nil {
				return nil, err
			}
			result = append(result, t)
		}
		return result, nil
	}

	switch r.Method {
	case "GET":
		panic("wow")
	case "PUT":
		panic("wow")
	case "POST":
		panic("wow")
	case "DELETE":
		panic("wow")
	}
}

func main() {
	http.HandleFunc(pathPrefix, apiHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
