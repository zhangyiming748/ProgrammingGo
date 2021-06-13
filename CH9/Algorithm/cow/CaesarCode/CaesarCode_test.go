package CaesarCode

import (
	"ProgrammingGo/log"
	"testing"
)

/*
一段由凯撒密码加密过的密文
凯撒密码指的是将字符偏移一定的单位
例如若偏移量为2，则a替换为c，b替换为d，c替换为e，...，z替换为b
若加密nowcoder，则密文为pqyeqfgt
现在发现加密包括数字、大写字母、小写字母，即0-9、A-Z、a-z的排列顺序进行偏移
现在截获了对方的一段密文以及偏移量，给定一段密文str和偏移量d，求对应的明文。
*/
func encode(str string, d int) string {
	// write code here
	offset := byte(d)
	log.Info.Println(offset)
	b := []byte(str)
	passwd := []byte{}
	for i := 0; i < len(b); i++ {
		if 'a' <= b[i] && b[i] <= 'z' {
			if b[i]+offset <= 'z' {
				passwd = append(passwd, b[i]+offset)
			} else {
				passwd = append(passwd, b[i]+offset-26)
			}
		}
		if 'A' <= b[i] && b[i] <= 'Z' {
			if b[i]+offset <= 'Z' {
				passwd = append(passwd, b[i]+offset)
			} else {
				passwd = append(passwd, b[i]+offset-26)
			}
		}
		if '0' <= b[i] && b[i] <= '9' {
			if b[i]+offset <= '9' {
				passwd = append(passwd, b[i]+offset)
			} else {
				passwd = append(passwd, b[i]+offset-10)
			}
		}

	}
	ret := string(passwd)
	return ret
}
func decode(str string, d int) string {
	// write code here
	offset := byte(d)
	log.Info.Println(offset)
	b := []byte(str)
	passwd := []byte{}
	for i := 0; i < len(b); i++ {
		if 'a' <= b[i] && b[i] <= 'z' {
			if b[i]-offset < 97 {

			}
		}
		if 'A' <= b[i] && b[i] <= 'Z' {
			offset = offset % 26
			if b[i]-offset >= 'A' {
				passwd = append(passwd, b[i]-offset)
			} else {
				passwd = append(passwd, b[i]-offset+26)
			}
		}
		if '0' <= b[i] && b[i] <= '9' {
			offset = offset % 10
			if b[i]-offset >= '0' {
				passwd = append(passwd, b[i]-offset)
			} else {
				passwd = append(passwd, b[i]-offset+10)
			}
		}

	}
	ret := string(passwd)
	return ret
}
func TestDecode(t *testing.T) {
	ret1 := decode("pqyeqfgt", 2)
	ret2 := decode("123ABCabc", 3)
	t.Logf("ret1期望输出\"nowcoder\"实际输出\"%s\"\nret2期望输出\"yz0789XYZ\"实际输出\"%s", ret1, ret2)
}
func TestByte(t *testing.T) {
	s := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz" //97 122 48 57 65 90
	b := []byte(s)
	t.Log(b)
}
func TestConvert(t *testing.T) {
	var b uint8
	var i int = 32
	b = uint8(i)
	t.Logf("%T", b)
}
