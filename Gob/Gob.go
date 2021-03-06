// Gob.
// Gob is a special encoding format in Golang that can replace JSON format

package main
import (
	"fmt"
	"encoding/gob"
	"bytes"
)

func main() {
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	data := map[string]string{"N":"J"}
	if err := enc.Encode(data); err != nil {
		fmt.Println(err)
	}
	const width = 16
	for start := 0; start < len(b.Bytes()); start += width {
		end := start + width
		if end > len(b.Bytes()) {
			end = len(b.Bytes())
		}
		fmt.Printf("%x\n", b.Bytes()[start:end])
	}
	dec := gob.NewDecoder(&b)
	var restored map[string]string
	if err := dec.Decode(&restored); err != nil {
		fmt.Println(err)
	}
	fmt.Println(restored)
}
