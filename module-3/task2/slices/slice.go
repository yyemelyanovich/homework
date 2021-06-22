package slices

import (
	"sort"
)

func SortStringsInPlace(s []string) {
	sort.Strings(s)
}

func SortStringsPure(s []string) []string {
	s1 := make([]string, len(s))
	copy(s1, s)
	sort.Strings(s1)
	return s1
}

type User struct {
	FirstName string
	LastName  string
}

func SortUsersPure(s []User) []User {
	r := make([]User, len(s))
	copy(r, s)
	sort.Slice(r, func(i, j int) bool {
		return (r[i].FirstName + r[i].LastName) < (r[j].FirstName + r[j].LastName)
	})
	return r
}

func SortUsersPureDesc(s []User) []User {
	r := make([]User, len(s))
	copy(r, s)
	sort.Slice(r, func(i, j int) bool {
		return (r[i].FirstName + r[i].LastName) > (r[j].FirstName + r[j].LastName)
	})
	return r
}
