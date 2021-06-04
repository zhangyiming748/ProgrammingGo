package magicNum

import (
	"fmt"
	"testing"
)

/*
描述
将字符串数字中为偶数位的数字进行翻转，将翻转后的结果进行输出。
示例1
输入：
"1234"
复制
返回值：
"1432"
复制
说明：
第2、4位为偶数，所以将其翻转后，得到 1432
示例2
输入：
"12346"
复制
返回值：
"16342"
复制
说明：
第2、4、5位为偶数，所以将其翻转后，得到 16342
备注：
数字的长度<=10^7 且不包含数字0
*/
func TestMagicNum(t *testing.T) {
	ex1 := "1234"
	ex2 := "12346"
	ret1 := change(ex1)
	ret2 := change(ex2)

	t.Logf("ret1=%s\tret2=%s", ret1, ret2)
}
func change(number string) string {
	if len(number) == 1 {
		return number
	}
	b := []byte(number)
	i := 0
	j := len(b) - 1
	fmt.Println(j)
	for {
		if i >= j {
			break
		}
		//i<j
		//fmt.Println(b[i])
		//fmt.Println(b[j])
		if b[i] == 48 || b[i] == 50 || b[i] == 52 || b[i] == 54 || b[i] == 56 {
			if b[j] == 48 || b[j] == 50 || b[j] == 52 || b[j] == 54 || b[j] == 56 { //i偶数j偶数
				b[i], b[j] = b[j], b[i]
				i++
				j--
			} else { //i偶数j奇数
				j--
			}

		}
		if b[i] == 49 || b[i] == 51 || b[i] == 53 || b[i] == 55 || b[i] == 57 {
			if b[j] == 48 || b[j] == 50 || b[j] == 52 || b[j] == 54 || b[j] == 56 { //i奇数j偶数
				i++
			} else { //i奇数j奇数
				i++
				j--
			}
		}

	}
	ret := string(b)
	return ret
}
func TestByte(t *testing.T) {
	num := "0123456789"
	//[48 49 50 51 52 53 54 55 56 57]
	b := []byte(num)
	fmt.Println(b)
}
