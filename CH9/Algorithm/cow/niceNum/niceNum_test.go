/*
描述
若一个数的首位和末位相等，则定义这个数为“好数”。
例如：1231、4512394是好数，而12345、768740则不是好数。
请你编写一个函数，判断是不是好数。如果是好数则返回true，否则返回false。
示例1
输入：
1231
复制
返回值：
true
复制
说明：
首位和末位都是1，相等。
示例2
输入：
4
复制
返回值：
true
复制
说明：
首位和末位都是4，相等。
示例3
输入：
100
复制
返回值：
false
复制
说明：
首位是1，末位是0，不相等。
备注：
1≤x≤10^91≤x≤10
9
*/
package niceNum

import (
	"fmt"
	"strconv"
	"testing"
)

func TestJudge(t *testing.T) {
	num := 12321
	ret := judge(num)
	t.Log(ret)
}
func judge(x int) bool {
	i := getFirst(x)
	fmt.Printf("i==%d", i)
	j := getLast(x)
	fmt.Printf("j==%d", j)
	if i != j {
		return false
	}
	return true
}
func getFirst(x int) int {
	s := strconv.Itoa(x)
	fmt.Println(s[0])
	xx := string(s[0])
	xxx, _ := strconv.Atoi(xx)
	return xxx

}
func getLast(x int) int {
	s := strconv.Itoa(x)
	fmt.Println(s[len(s)-1])
	xx := string(s[len(s)-1])
	xxx, _ := strconv.Atoi(xx)
	return xxx
}
func TestGetLast(t *testing.T) {
	num := 12345
	ret := getLast(num)
	t.Log(ret)
}
