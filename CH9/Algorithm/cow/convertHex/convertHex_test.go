package convertHex

import (
	"strings"
	"testing"
)

/*
描述
给定一个十进制数M，以及需要转换的进制数N。将十进制数M转化为N进制数
示例1
输入：
7,2
复制
返回值：
"111"

*/
/*
string solve(int M, int N) {
        string res;
        bool flag = false;
        if (M < 0) {
            M = -M;
            flag = true;
        }
        const char* a = "0123456789ABCDEF";
        while (M != 0) {
            res = a[M % N] + res;
            M /= N;
        }
        if (flag) {
            res = "-" + res;
        }
        return res;
    }
*/
const a = "0123456789ABCDEF"

func solve(M int, N int) string {
	var res []byte
	var flag bool
	for M != 0 {
		res = append(res, a[M%N])
		M /= N
	}
	result := string(res)
	if flag {
		strings.Join([]string{"-", result}, "")
	}
	return result
}
func TestSolve(t *testing.T) {
	ret := solve(7, 2)
	t.Log(ret)
}
