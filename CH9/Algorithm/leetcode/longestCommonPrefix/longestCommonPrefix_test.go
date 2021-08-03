package longestCommonPrefix

import (
	"math"
	"testing"
)

func longestCommonPrefix(strs []string) string {
	i, s := stand(strs)
	retNum := same(i, s, strs)
	sb := []byte(s)
	ret := sb[:retNum]
	return string(ret)
}
func stand(strs []string) (int, string) {
	short := math.MaxInt32
	shortstr := ""
	for _, v := range strs {
		if l := len(v); l < short {
			short = l
			shortstr = v
		}
	}
	return short, shortstr
}
func same(l int, s string, strs []string) int {
	count := 0
	for i := 0; i < l; i++ {
		for _, v := range strs {
			if v[i] == s[i] {
				continue
			} else {
				return count
			}
		}
		count++
	}
	return l
}
func TestUnit(t *testing.T) {
	strs := []string{"flower", "flow", "flight"}
	//"fl"
	ret := longestCommonPrefix(strs)
	t.Log(ret)
}
