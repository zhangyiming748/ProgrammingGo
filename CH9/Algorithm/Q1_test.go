package main

import (
	"sort"
	"testing"
)

func TestQ1(t *testing.T) {
	input := []int{5, 100, 100, 100, 120, 130, 40, 40, 30, 60, 50}
	ret := Q1(input)
	t.Log(ret)
	sort.Ints(input)
}
