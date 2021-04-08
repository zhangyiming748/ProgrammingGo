package CH1

import (
	"fmt"
	"log"
)

var bigDigits = [][]string{
	{"  000  "},
	{" 0   0 "},
	{"0     0"},
	{"0     0"},
	{"0     0"},
	{" 0   0 "},
	{"  000  "}, {""},
}

func BigNum(num string) {

	for row := range bigDigits[0] {
		line := ""
		for column := range num {
			digit := num[column] - '0'
			if 0 <= digit && digit <= 9 {
				line += bigDigits[digit][row] + " "
			} else {
				log.Fatal("error")
			}
		}
		fmt.Println(line)
	}
}
