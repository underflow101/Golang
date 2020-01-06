// fundamental sort

package main
import (
	"fmt"
	"sort"
)

func main() {
	// int slice sort
	s := []int{20, 12, 23, 1, 2, 5, 4, 3, 7, 2, 4, 6, 8}
	d := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	sort.Sort(sort.IntSlice(s))					// int slice sort in ascending order
	sort.Sort(sort.Reverse(sort.IntSlice(d)))	// int slice sort in descending order
	fmt.Println(s)
	fmt.Println(d)

	// float64 slice sort
	a := []float64{1.2, 2.3215, 7.231245, 10.2222, 10,2223, 10.2224, 12}
	b := []float64{10, 9, 2.2, 11.2, 62.2}
	sort.Sort(sort.Reverse(sort.Float64Slice(a)))	// float64 sort in descending order
	sort.Sort(sort.Float64Slice(b))					// float64 sort in ascending order
	fmt.Println(a)
	fmt.Println(b)

	// string slice sort
	q := []string{"Work", "work", "job", "rest", "life", "birth", "Death"}
	sort.Sort(sort.StringSlice(q))
	fmt.Println(q)
	sort.Sort(sort.Reverse(sort.StringSlice(q)))
	fmt.Println(q)
}