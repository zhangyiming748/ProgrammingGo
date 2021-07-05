package MyAtoi

/*

 */
import (
	"testing"
)

func TestByte(t *testing.T) {
	nums := "a z A Z"
	b := []byte(nums)
	t.Log(b)
}
func TestAtoi(t *testing.T) {
	in := "      -117e40091539"

	ret := atoi(in)
	t.Log("success", ret)

}
func atoi(str string) int {
	b := []byte(str)
	num := 0
	nagitave := false

	for _, v := range b {
		if ASCII(v) {
			break
		}
		if illegal(v) {
			continue
		}

		if v == 32 || v == 43 {
			continue
		}
		if v == 45 {
			nagitave = true
			continue
		}

		switch v {
		case 48:
			num = num*10 + 0
		case 49:
			num = num*10 + 1
		case 50:
			num = num*10 + 2
		case 51:
			num = num*10 + 3
		case 52:
			num = num*10 + 4
		case 53:
			num = num*10 + 5
		case 54:
			num = num*10 + 6
		case 55:
			num = num*10 + 7
		case 56:
			num = num*10 + 8
		case 57:
			num = num*10 + 9
		}
	}
	if nagitave {
		num = 0 - num

	}
	if num > 2147483647 {
		num = 2147483647
	}
	if num < -2147483648 {
		num = -2147483648
	}
	return num
}

func illegal(b byte) bool {
	if b >= 48 && b <= 57 {
		return false
	}
	if b == 43 || b == 45 {
		return false
	}
	return true
}
func ASCII(b byte) bool {
	if b >= 97 && b <= 122 || b >= 65 && b <= 90 {
		return true
	}
	return false
}
