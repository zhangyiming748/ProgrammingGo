package lengthOfLongestSubstring

import "testing"

func lengthOfLongestSubstring(s string) int {
<<<<<<< HEAD
	long := 0

	for i := 0; i < len(b); i++ {
		for j := 1; j < len(b); j++ {
			if s[i] == s[j] {
=======
	sb := []byte(s)
	longest := 0
	for i := 0; i < len(sb); i++ {
		long := 0
		for j := i + 1; j < len(sb); j++ {
			if sb[i] == sb[j] {
				if long > longest {
					long = longest
				}
>>>>>>> f5356df498b5673f0aa0fbb5ac065ff1ebd82d79
				break
			}
			long++
		}
	}

	return longest
}
func TestSubstring(t *testing.T) {
	s := "abcabcbb"
	ret := lengthOfLongestSubstring(s)
	t.Log(ret)
}
