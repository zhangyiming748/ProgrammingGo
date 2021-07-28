package lengthOfLongestSubstring

import "testing"

func lengthOfLongestSubstring(s string) int {
	sb := []byte(s)
	longest := 0
	for i := 0; i < len(sb); i++ {
		long := 0
		for j := i + 1; j < len(sb); j++ {
			if sb[i] == sb[j] {
				if long > longest {
					long = longest
				}
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
