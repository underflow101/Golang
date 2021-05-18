// JSON field Manipulation

package main
import (
	"fmt"
	"encoding/json"
)

type Fields struct {
	VisibleField string
	InvisibleField string
}

func main() {
	f := &Fields{"a", "b"}
	b, _ := json.Marshal(struct {
		*Fields
		InvisibleField string `json:",omitempty"`
		Additional string
	}{Fields: f, Additional: "c"})
	fmt.Println(string(b))
}
