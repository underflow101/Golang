// BOJ Golang Solve
// 15552 - 빠른 A+B

package main

import (
	"os"
	"bufio"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	scanner.Scan()
	T, _ := strconv.Atoi(scanner.Text())

	for i := T; i > 0; i-- {
		scanner.Scan()
		items := strings.Split(scanner.Text(), " ")
		sum := 0
		for _, item := range items {
			newitem, _ := strconv.Atoi(item)
			sum += newitem
		}
		out.WriteString(strconv.Itoa(sum) + "\n")
	}
	out.Flush()
}
