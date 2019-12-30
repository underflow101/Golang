// Filestream basic

package main

import (
	"fmt"
	"os"
)

func main() {
	dir := "/users/bear.paek/desktop/fs.txt"
	f, err := os.Open(dir)
	if err != nil {
		fmt.Println("NO! ERROR!")
		//return err
	}

	defer f.Close()
	var num int
	if _, err := fmt.Fscanf(f, "%d\n", &num); err == nil {
		fmt.Println("YES!!")
		fmt.Println(num)
	}
}
