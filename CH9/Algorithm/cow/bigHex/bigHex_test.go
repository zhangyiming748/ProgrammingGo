package bigHex

import (
	"fmt"

	"strconv"
	"testing"
)

/*
描述
给定一个包含大写英文字母和数字的句子，找出这个句子所包含的最大的十六进制整数，返回这个整数的值。数据保证该整数在int表示范围内

示例1
输入：
"012345BZ16"
复制
返回值：
1193051
复制
说明：
12345B对应十进制为1193051
*/
func TestBigHex(t *testing.T) {
	ex := "012345BZ16"
	ret := solve(ex)
	t.Log(ret)
}
func BenchmarkBigHex(b *testing.B) {
	ex := "012345BZ16"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		solve(ex)
	}
	b.StopTimer()
}
func solve(s string) int {
	var res uint64 = 0
	max := 0
	for i := 0; i < len(s); i++ {
		if illegal(s[i]) {
			continue
		}
		for j := i + 1; j < len(s); j++ {
			if illegal(s[j]) {
				tmp := s[i:j]
				res, _ = strconv.ParseUint(tmp, 16, 32)
				if int(res) > max {
					max = int(res)
				}
			}
		}
	}

	return max
}
func wrongSolve(s string) int {
	var all []string
	b := []byte(s)
	for i := 0; i < len(b)-1; i++ {
		if illegal(b[i]) {
			continue
		}
		for j := i + 1; j < len(b)-1; j++ {
			if illegal(b[j]) {
				all = append(all, string(b[i:j]))
			}
		}

	}
	bigdex := chooesBiggest(all)
	return bigdex

}
func chooesBiggest(s []string) int {
	var bigdex uint64 = 0
	for _, v := range s {
		dex, _ := strconv.ParseUint(v, 16, 32)
		if dex > bigdex {
			bigdex = dex
		}
	}
	return int(bigdex)
}

func TestHex(t *testing.T) {
	dict := "0123456789ABCDEF"
	for i, v := range dict {
		fmt.Printf("index%d\tvalve is %v\n", i, v)
	}

}
func illegal(i byte) bool {
	if i == 48 || i == 49 || i == 50 || i == 51 || i == 52 || i == 53 || i == 54 || i == 55 || i == 56 || i == 57 || i == 65 || i == 66 || i == 67 || i == 68 || i == 69 || i == 70 {
		return false
	}
	return true
}
