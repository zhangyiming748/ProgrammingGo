/*
描述
给定一个数组由一些非负整数组成，现需要将他们进行排列并拼接，使得最后的结果最大，返回值需要是string类型 否则可能会溢出
示例1
输入：
[30,1]
复制
返回值：
"301"
复制
备注：
输出结果可能非常大，所以你需要返回一个字符串而不是整数。
关联企业
*/
package biggest

import (
	"ProgrammingGo/log"
	"fmt"
	"sort"
	"strconv"
	"testing"
)

func solve(nums []int) string {
	l := len(nums)
	zero := 0
	for _, v := range nums {
		if v == 0 {
			zero++
		}
	}
	if zero == l {
		return "0"
	}
	single := []int{}
	for _, v := range nums {
		for {
			if v <= 0 {
				break
			}
			number := v % 10
			v = v / 10
			single = append(single, number)
		}
	}
	log.Info.Println(single)
	ret := beforeSolve(single)
	return ret
}
func beforeSolve(nums []int) string {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	str := ""
	for i, v := range nums {
		sv := strconv.Itoa(v)
		str = fmt.Sprintf("%s%s", str, sv)
		log.Info.Printf("拼接后的字符串%s\n", str)
		log.Info.Printf("拼接第%v个数字%v\n", i, v)
	}
	log.Info.Println(str)
	return str
}
func TestSolve(t *testing.T) {
	in := []int{10, 1}
	ret := solve(in)
	t.Log(ret)
}
