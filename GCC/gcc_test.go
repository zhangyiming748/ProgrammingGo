package GCC

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestGcc(t *testing.T) {
	gcc()
}

/*
标题：按身高和体重排队 时间限制：1秒 内存限制：262144K 语言：不限

某学校举行运动会，学生们按编号1、2、3、....n进行标识，现需要按照身高由低到高排列，对身高相同的人，按体重由轻到重排列；对身高和体重都相同的人，维持原有的编号顺序关系。请输出排列后的学生编号。

输入描述：两个序列，每个序列由n个正整数组成0<n<=100。第一个序列中的数值代表身高，第二个序列中的数值代表体重。
输出描述：排列结果，每个数值都是原始序列中的学生编号，编号从1开始

比如：
输入：4 100 100 120 130 40 30 60 50
输出：2 1 3 4
*/
func TestQ1(t *testing.T) {

}

/*
标题：找出符合要求的字符串子串 时间限制：1秒 内存限制：262144K 语言：不限

给定两个字符串，从字符串2中找出字符串1中的所有字符，去重并按照ASCII值从小到大排序
输入字符串1:长度不超过1024
输入字符串2:长度不超过1000000
字符串范围满足ASCII编码要求，按照ASCII的值由小到大排序。

输入描述：
bach
bbaaccedfg
输出描述：
abc

备注：输入字符串1为给定字符串bach，输入字符串2  bbaaccedfg ，从字符串2中找出字符串1的字符，去除重复的字符，并且按照ASCII值从小到大排序，得到输出的结果为abc
字符串1中的字符h在字符串2中找不到不输出

示例：
输入：
fach
bbaaccedfg
输出：
acf
*/
func TestQ2(t *testing.T) {
	s1 := "fach"
	s2 := "bbaaccedfg"
	ans := ""
	for _, v := range s1 {
		if strings.Contains(s2, string(v)) {
			if !strings.Contains(ans, string(v)) {
				ans += string(v)
			}
		}
	}
	t.Logf("ans =%s \n", ans)
	ret := MySort(ans)
	t.Logf("ret = %v\n", ret)

}
func MySort(s string) string {
	r := make([]rune, 0)

	for _, v := range s {
		r = append(r, v)
	}
	for i := 0; i < len(r); i++ {
		for j := 1; j < len(r)-i; j++ {
			if r[j] < r[j-1] {
				//交换
				r[j], r[j-1] = r[j-1], r[j]
			}
		}
	}
	var ret string
	for _, val := range r {
		ret = ret + string(val)
	}
	return ret
}

/*
找最小数
标题：找最小数 时间限制：1秒 内存限制：32768K 语言：不限
给一个正整数NUM1，计算出新正整数NUM2，NUM2为NUM1中移除N位数字后的结果，需要使得NUM2的值最小。

输入描述：1.输入的第一行为一个字符串，字符串由0～9字符组成，记录正整数NUM1，NUM1长度小于32。2.输入的第二行为需要移除的数字的个数，小于NUM1长度。
输出描述：输出一个数字字符串，记录最小值NUM2

示例：
输入：
2615371
4
输出：
131
*/
func TestMix(t *testing.T) {
	numbers := 2695371
	snum := strconv.Itoa(numbers)
	newNumbers := new(string)
	l := len(snum)
	del := 4
	for i := 0; i <del; i++ {
		if strings.Contains(snum, "9") {
			index := strings.Index(snum, "9")
			if index > l-1 {
				index = l - 1
			}
			*newNumbers = snum[:index] + snum[index+1:]
		} else if strings.Contains(snum, "8") {
			index := strings.Index(snum, "8")
			if index > l-1 {
				index = l - 1
			}
			*newNumbers = snum[:index] + snum[index+1:]
		} else if strings.Contains(snum, "7") {
			index := strings.Index(snum, "7")
			if index > l-1 {
				index = l - 1
			}
			*newNumbers = snum[:index] + snum[index+1:]
		} else if strings.Contains(snum, "6") {
			index := strings.Index(snum, "6")
			if index > l-1 {
				index = l - 1
			}
			*newNumbers = snum[:index] + snum[index+1:]
		} else if strings.Contains(snum, "5") {
			index := strings.Index(snum, "5")
			if index > l-1 {
				index = l - 1
			}
			*newNumbers = snum[:index] + snum[index+1:]
		} else if strings.Contains(snum, "4") {
			index := strings.Index(snum, "4")
			if index > l-1 {
				index = l - 1
			}
			*newNumbers = snum[:index] + snum[index+1:]
		} else if strings.Contains(snum, "3") {
			index := strings.Index(snum, "3")
			if index > l-1 {
				index = l - 1
			}
			*newNumbers = snum[:index] + snum[index+1:]
		} else if strings.Contains(snum, "2") {
			index := strings.Index(snum, "2")
			if index > l-1 {
				index = l - 1
			}
			*newNumbers = snum[:index] + snum[index+1:]
		} else if strings.Contains(snum, "1") {
			index := strings.Index(snum, "1")
			if index > l-1 {
				index = l - 1
			}
			*newNumbers = snum[:index] + snum[index+1:]
		}
	}
	ret, _ := strconv.Atoi(*newNumbers)
	fmt.Println(ret)
}
