package cow

import (
	. "ProgrammingGo/log"
	"math"
	"time"
)

func longestCommonPrefix(strs []string) string {
	min:=getShort(strs)
	for _,v:=range strs{
		for i:=0;i<min;i++{
			if strs[v][i]
		}
	}
	return ""
}

func getShort(s []string) int {
	length := math.MaxInt32
	for _, v := range s {
		Debug.Printf("此次循环字符串长度%d", len(v))
		if len(v) < length {
			length = len(v)
		}
	}
	return length
}
