// JSON tag & User customized JSON serialization

package main

import (
	"errors"
	"fmt"
	"time"
)

type status int
type Deadline struct {
	time.Time
}

const (
	UNKNOWN status = iota
	TODO
	DONE
	YET
)

func NewDeadline(t time.Time) *Deadline {
	return &Deadline{t}
}

type Task struct {
	Title    string
	Status   status
	Deadline *Deadline
}

// JSON tag
type Myjson struct {
	Title    string `json:"title"`     // Title 대신 title을 JSON 필드로 사용
	Internal string `json:"-"`         // 해당 필드가 JSON에 나타나지 않고 무시됨
	Value    int64  `json:",omitempty` // 0인 경우 JSON에서 결과를 내지 않음
	ID       int64  `json:",string"`   // JSON에서는 문자열로 형변환됨
	// 자바스크립트에서는 정수값이 53비트를 넘어서면 정확도가 떨어지므로,
	// 64비트 정수를 JSON으로 주고받을 때에는 형변환을 습관화할 것
}

// User Customized JSON Serialization
func (s status) MarshalJSON() ([]byte, error) {
	switch s {
	case UNKNOWN:
		return []byte(`"UNKNOWN"`), nil
	case TODO:
		return []byte(`"TODO"`), nil
	case DONE:
		return []byte(`"DONE"`), nil
	default:
		return nil, errors.New("status.MarshalJSON: unknown value")
	}
}

// User Customized JSON Unserialization
func (s *status) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"UNKNOWN"`:
		*s = UNKNOWN
	case `"TODO"`:
		*s = TODO
	case `"DONE"`:
		*s = DONE
	default:
		return errors.New("status.UnmarshalJSON: unknown value")
	}
	return nil
}

func main() {
	b := []byte(`{"Title":"Buy Milk","Status":2,"Deadline":"2020-05-26T23:59:59Z"}`)
	t := Task{}
	t.Status.UnmarshalJSON(b)
	fmt.Println(t)
}
