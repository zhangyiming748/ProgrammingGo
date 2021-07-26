package lengthOfLongestSubstring

func lengthOfLongestSubstring(s string) int {
	long := 0
	b := []byte(s)
	for i := 0; i < len(b); i++ {
		for j := 1; j < len(b); j++ {
			if b[i] == b[j] {
				break
			} else {
				long++
			}
		}
	}
}
