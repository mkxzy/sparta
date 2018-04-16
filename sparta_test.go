package main

import "testing"

func TestSlice(t *testing.T) {
	s := make([]int, 0, 1)
	s = append(s, 1)
	s = append(s, 2)
	t.Log(len(s), cap(s))
	t.Log(s)
}
