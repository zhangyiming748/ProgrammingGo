package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"testing"
)

func TestOutput(t *testing.T) {
	t.Logf("%.12f", rand.Float64())
	s1 := fmt.Sprintf("%.12f", rand.Float64())
	radio, _ := strconv.Atoi(strings.Split(fmt.Sprintf("%.12f", rand.Float64()), ".")[1])
	t.Logf("radio = %v\n", radio)
	t.Logf("s1 = %v\n", s1)
}
