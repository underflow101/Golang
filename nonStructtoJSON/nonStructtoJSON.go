// Non-struct to JSON

package main
import (
	"fmt"
	"encoding/json"
)

func main() {
	// map[string]string
	b, _ := json.Marshal(map[string]string {
		"Name": "John",
		"Age": "16",
	})
	fmt.Println(string(b))

	// map[string]interface
	c, _ := json.Marshal(map[string]interface{} {
		"Name":"John",
		"Age": 16,
	})
	fmt.Println(string(c))
}