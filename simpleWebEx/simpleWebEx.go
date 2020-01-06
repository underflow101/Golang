// simple & basic web example in golang

package main
import (
	"log"
	"net/http"
	"fmt"
)

func main() {
	fmt.Println("Start!!!")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Println("end")
}