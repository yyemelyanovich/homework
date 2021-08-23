package naming

import (
	"testing"
)

var result int

func BenchmarkGetNumberSplit(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = getNumberSplit("12.11")
	}
	result = r
}

func BenchmarkGetNumberRegexp(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = getNumberRegexp("12.11")
	}
	result = r
}
