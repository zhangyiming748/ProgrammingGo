package cow

import . "ProgrammingGo/log"

//使用前缀后缀匹配方法
func judge(str string) bool {
	prefix := ""
	suffix := ""
	mid := 0
	if len(str) == 1 {
		return true
	}
	if len(str) == 2 {
		if str[0] == str[1] {
			return true
		} else {
			return false
		}
	}
	if len(str)%2 == 0 {
		mid = len(str) / 2
		prefix = str[0:mid]
		suffix = str[mid:]
	} else {
		mid = len(str) / 2
		prefix = str[0:mid]
		suffix = str[mid+1:]
	}
	antiSuffix := antiString(suffix)
	if prefix == antiSuffix {
		return true
	}
	Debug.Printf("mid = %v\tprefix =%v\tsuffix =%v", mid, prefix, suffix)
	return false

}
func antiString(s string) string {
	b := []byte(s)
	l := len(b)
	nline := []byte{}
	for i := 0; i < l; i++ {
		nline = append(nline, b[l-i-1])
	}
	return string(nline)
}
