package stringmultiset

import "strings"

// MultiSet is a data type idealisation
type MultiSet map[string]int

//Insert Func
func (m MultiSet) Insert(val string) {
	m[val]++
}

// Erase func
func (m MultiSet) Erase(val string) {
	if m[val] < 2 {
		delete(m, val)
	} else {
		m[val]--
	}
}

// Count func
func (m MultiSet) Count(val string) int {
	return m[val]
}

func (m MultiSet) String() string {
	s := "{ "
	for val, count := range m {
		s += strings.Repeat(val+" ", count)
	}
	return s + " }"
}
