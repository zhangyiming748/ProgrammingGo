package CaesarCode

import "testing"

/*
一段由凯撒密码加密过的密文
凯撒密码指的是将字符偏移一定的单位
例如若偏移量为2，则a替换为c，b替换为d，c替换为e，...，z替换为b
若加密nowcoder，则密文为pqyeqfgt
现在发现加密包括数字、大写字母、小写字母，即0-9、A-Z、a-z的排列顺序进行偏移
现在截获了对方的一段密文以及偏移量，给定一段密文str和偏移量d，求对应的明文。
*/
func decode(str string, d int) string {
	// write code here
	offset := byte(d)
	b := []byte(str)
	passwd := []byte{}
	for i := 0; i < len(b); i++ {
		if 'a' <= b[i] && b[i] <= 'z' {
			if newer := b[i] + offset; newer > 'z' {
				passwd = append(passwd, newer-'z'+'a')
			} else {
				passwd = append(passwd, newer)
			}
		}
		if 'A' <= b[i] && b[i] <= 'Z' {
			if newer := b[i] + offset; newer > 'Z' {
				passwd = append(passwd, newer-'Z'+'A')
			} else {
				passwd = append(passwd, newer)
			}
		}
		if '0' <= b[i] && b[i] <= '9' {
			if newer := b[i] + offset; newer > '9' {
				passwd = append(passwd, newer-'9'+'0')
			} else {
				passwd = append(passwd, newer)
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
	s := "az09AZ0123456789" //97 122 48 57 65 90
	b := []byte(s)
	t.Log(b)
}
