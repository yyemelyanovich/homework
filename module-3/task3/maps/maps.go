package maps

import (
	"sort"
)

func SortedValues(m map[int]string) string {
	var out string
	s := make([]int, 0, len(m))
	for i, _ := range m {
		s = append(s, i)
	}
	sort.Ints(s)
	for _, i := range s {
		out = out + m[i]
	}
	return out
}
