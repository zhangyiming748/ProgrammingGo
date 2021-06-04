package moveA

import "testing"

/*
移动字母
 算法知识视频讲解

知识点
字符串
题目
题解(4)
讨论(14)
排行
描述
给定一个只包含小写字母的字符串s，牛牛想将这个字符串中的所有'a'字母全部移动到字符串的末尾，而且保证其它字符的相对顺序不变。其中字符串s的长度<=1e6。
示例1
输入：
"abcavv"
复制
返回值：
"bcvvaa"
*/
func change(s string) string {
	b := []byte(s)
	nb := []byte{}
	count := 0
	for i := 0; i < len(b); i++ {
		if b[i] == 97 {
			count++
		} else {
			nb = append(nb, b[i])
		}
	}
	for j := 0; j < count; j++ {
		nb = append(nb, 97)
	}
	sb := string(nb)
	return sb
}
func TestChange(t *testing.T) {
	s := "abcavv"
	ret := change(s)
	t.Log(ret)
}
func TestByte(t *testing.T) {
	s := "abc"
	b := []byte(s)
	t.Log(b)
}
