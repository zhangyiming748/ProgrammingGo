package cow

import (
	. "ProgrammingGo/log"
	"math"
	//"time"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	min, word := getShort(strs)
	if min == 1 {
		//tmp:=strs[0][0]
		//for i,v:=range strs {
		//	if tmp == []byte(v)/ {
		//		tmp = v
		//	}else {
		//		return ""
		//	}
		//
		//}
		return strs[0]
	}
	count := 0
	for _, v := range strs { //遍历每个单词
		if count < min {
			if v[count] != word[count] {
				Debug.Println(count)
				break

			} else {
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
