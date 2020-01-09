// ParallelMin!

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func Min(a []int) int {
	if len(a) == 0 {
		return 0
	}
	min := a[0]
	for _, e := range a[1:] {
		if min > e {
			min = e
		}
	}
	return min
}

func ParallelMin(a []int, n int) int {
	if len(a) < n {
		return Min(a)
	}
	mins := make([]int, n)
	size := (len(a) + n - 1) / n
	var wait sync.WaitGroup
	for i := 0; i < n; i++ {
		wait.Add(1)
		go func(i int) {
			defer wait.Done()
			begin, end := i*size, (i+1)*size
			if end > len(a) {
				end = len(a)
			}
			mins[i] = Min(a[begin:end])
		}(i)
	}
	wait.Wait()
	return Min(mins)
}

func main() {
	runtime.GOMAXPROCS(4)
	fmt.Println(ParallelMin([]int{
		83, 99, 231, 2155, 3646,
		2332, 2, 1, 667, 234, 795,
		3357, 57787, 32, 3233, 333333333,
		9, -2, -5, -20335, -1,
	}, 5))
}
