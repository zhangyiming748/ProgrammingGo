package cow

import (
	. "ProgrammingGo/log"
	"math"
	//"time"
)

func longestCommonPrefix(strs []string) string {
	min, word := getShort(strs)

	count := 0
	for i := 0; i < len(strs); i++ { //遍历每个单词
		for j := 1; j < min; j++ {
			before := strs[i][j-1]
			after := strs[i][j]
			if before == after {
				count++
			}
		}
	}
	Debug.Println(count)
	return word[:count]
}

func getShort(s []string) (int, string) {
	length := math.MaxInt32
	str := ""
	for _, v := range s {
		Debug.Printf("此次循环字符串长度%d", len(v))
		if len(v) < length {
			length = len(v)
			str = v
		}
	}
	return length, str
}
