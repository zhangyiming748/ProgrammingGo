package verifySTR

import (
	"strings"
	"testing"
)

func illegal(ch byte) bool {
	if (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9') {
		return false
	}
	return true
}
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	start := 0
	end := len(s) - 1
	for start < end {
		for start < end && illegal(s[start]) {
			start++
		}
		for start < end && illegal(s[end]) {
			end--
		}
		if start < end {
			if s[start] != s[end] {
				return false
			}
			start++
			end--
		}
	}
	return true
}
func TestUnit(t *testing.T) {
	str1 := "A man, a plan, a canal: Panama"
	str2 := "race a car"
	t.Logf("str1 is %v,str2 is %v", isPalindrome(str1), isPalindrome(str2))
}
