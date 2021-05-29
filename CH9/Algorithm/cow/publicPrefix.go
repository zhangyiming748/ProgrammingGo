package cow

import (
	"log"
	"math"
	"strconv"
	"strings"
)

func publicPrefix(s []string) string {
	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		return s[0]
	}
	short := shorter(s) //最短字符串
	var count int       //每一个字符串的内部索引
	var sameword []byte
	for count = 0; count < short; count++ {
		word := []byte{}
		for _, v := range s {
			word = append(word, v[count])
			log.Println("word = ", word)
		}
		if same(word) {
			count += 1
			//sameword = append(sameword, uint8(word))

		} else {
			break
		}
	}
	log.Println(sameword)
	ret := strconv.Itoa(count)
	return string(ret)
}
func same(b []byte) bool {
	for i, v := range b {
		if i > 1 {
			if b[i-1] == v {
				log.Println("相同")
				continue
			} else {
				return false
			}
		}
	}
	return true
}
func shorter(s []string) int {
	short := math.MaxInt32
	for _, v := range s {
		s := len(v)
		if s < short {
			short = s
		}
	}
	return short
}
func hasprefix(s, p string) bool {
	if strings.HasPrefix(s, p) {
		return true
	} else {
		return false
	}
}
