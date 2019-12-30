// Filestream 2

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func writeTo(w io.Writer, lines []string) error {
	for _, line := range lines {
		if _, err := fmt.Fprintln(w, line); err != nil {
			return err
		}
	}
	return nil
}

func readFrom(r io.Reader, lines *[]string) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		*lines = append(*lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func exampleWriteTo() {
	lines := []string{
		"bill@mail.com",
		"tom@mail.com",
		"jane@mail.com",
		"great@mail.com",
		"Hello",
		"GOOOOD",
	}
	if err := writeTo(os.Stdout, lines); err != nil {
		fmt.Println(err)
	}
}

func exampleReadFrom() {
	r := strings.NewReader("bill\ntom\njane\n")
	var lines []string
	if err := readFrom(r, &lines); err != nil {
		fmt.Println(err)
	}
	fmt.Println(lines)
}

func main() {
	exampleWriteTo()
	exampleReadFrom()
}
