package main

import (
	"strconv"
	"strings"
)

type Numbers struct {
	TypeInt    int
	TypeString string
}

func (N *Numbers) SetNumByString(numbers string) {

	N.TypeString = numbers
	N.TypeInt, _ = strconv.Atoi(numbers)
	return
}
func (N *Numbers) SetNumByInt(numbers int) {
	N.TypeInt = numbers
	N.TypeString = strconv.Itoa(numbers)
	return
}
func (N Numbers) GetString() string {
	return N.TypeString
}
func (N Numbers) GetInt() int {
	return N.TypeInt
}
func (N Numbers) GetLen() int {
	l := len(N.TypeString)
	return l
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
func FindMin(num, del int) int {
	var n Numbers
	n.SetNumByInt(num)
	//del := 4
	l := n.GetLen()
	for i := 0; i < del; i++ {
		if strings.Contains(n.TypeString, "9") {
			index := strings.Index(n.TypeString, "9")
			if index > l-1 {
				index = l - 1
			}
			ret := n.TypeString[:index] + n.TypeString[index+1:]
			n.SetNumByString(ret)
		} else if strings.Contains(n.TypeString, "8") {
			index := strings.Index(n.TypeString, "8")
			if index > l-1 {
				index = l - 1
			}
			ret := n.TypeString[:index] + n.TypeString[index+1:]
			n.SetNumByString(ret)
		} else if strings.Contains(n.TypeString, "7") {
			index := strings.Index(n.TypeString, "7")
			if index > l-1 {
				index = l - 1
			}
			ret := n.TypeString[:index] + n.TypeString[index+1:]
			n.SetNumByString(ret)
		} else if strings.Contains(n.TypeString, "6") {
			index := strings.Index(n.TypeString, "6")
			if index > l-1 {
				index = l - 1
			}
			ret := n.TypeString[:index] + n.TypeString[index+1:]
			n.SetNumByString(ret)
		} else if strings.Contains(n.TypeString, "5") {
			index := strings.Index(n.TypeString, "5")
			if index > l-1 {
				index = l - 1
			}
			ret := n.TypeString[:index] + n.TypeString[index+1:]
			n.SetNumByString(ret)
		} else if strings.Contains(n.TypeString, "4") {
			index := strings.Index(n.TypeString, "4")
			if index > l-1 {
				index = l - 1
			}
			ret := n.TypeString[:index] + n.TypeString[index+1:]
			n.SetNumByString(ret)
		} else if strings.Contains(n.TypeString, "3") {
			index := strings.Index(n.TypeString, "3")
			if index > l-1 {
				index = l - 1
			}
			ret := n.TypeString[:index] + n.TypeString[index+1:]
			n.SetNumByString(ret)
		} else if strings.Contains(n.TypeString, "2") {
			index := strings.Index(n.TypeString, "2")
			if index > l-1 {
				index = l - 1
			}
			ret := n.TypeString[:index] + n.TypeString[index+1:]
			n.SetNumByString(ret)
		} else if strings.Contains(n.TypeString, "1") {
			index := strings.Index(n.TypeString, "1")
			if index > l-1 {
				index = l - 1
			}
			ret := n.TypeString[:index] + n.TypeString[index+1:]
			n.SetNumByString(ret)
		} else if strings.Contains(n.TypeString, "0") {
			index := strings.Index(n.TypeString, "0")
			if index > l-1 {
				index = l - 1
			}
			ret := n.TypeString[:index] + n.TypeString[index+1:]
			n.SetNumByString(ret)
		}
	}
	return n.GetInt()
}
