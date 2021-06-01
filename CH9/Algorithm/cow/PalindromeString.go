package cow

import (
	. "ProgrammingGo/log"
)

func judge(str string) bool {
	b := []byte(str)
	mid := 0
	var l, r int
	if len(str)%2 == 0 {
		mid = len(str) / 2
	} else {
		mid = len(str)/2 + 1

		for mid-1 > 0 && mid+1 < len(b) {
			l = mid - 1
			Debug.Println(l)
			r = mid + 1
			Debug.Println(r)
			if b[l] != b[r] {
				return false
			}
		}
	}

	return true
}

func halfString(s string) string {
	b := []byte(s)
	mid := len(b) / 2
	return string(b[:mid])
}
