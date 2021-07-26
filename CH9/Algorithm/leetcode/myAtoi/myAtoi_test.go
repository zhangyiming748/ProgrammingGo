package myAtoi

import (
	"strconv"
	"testing"
)

func myAtoi(s string) int {
	b := []byte(s)
	var nagitive bool
	num:=0
	for _, v := range b {
		switch v {
		case '-':
			nagitive = true
		case '0':

		case '1':

		case '2':

		case '3':

		case '4':

		case '5':

		case '6':

		case '7':
		case '8':
		case '9':

		}
	}
}
func TestByte(t *testing.T) {
	s := "+-0123456789"
	b := []byte(s)
	t.Log(b)
}
