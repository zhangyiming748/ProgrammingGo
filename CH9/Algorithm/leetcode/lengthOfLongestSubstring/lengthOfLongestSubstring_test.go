package lengthOfLongestSubstring

func lengthOfLongestSubstring(s string) int {
	long := 0

	for i := 0; i < len(b); i++ {
		for j := 1; j < len(b); j++ {
			if s[i] == s[j] {
				break
			} else {
				long++
			}
		}
	}
}
