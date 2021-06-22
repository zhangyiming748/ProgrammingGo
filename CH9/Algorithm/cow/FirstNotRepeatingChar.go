package cow

func FirstNotRepeatingChar(str string) int {
	s := []byte(str)
	m := make(map[byte]int)
	for _, v := range s {
		if _, ok := m[v]; ok {
			m[v] = m[v] + 1
		} else {
			m[v] = 1
		}
	}
	for i, value := range s {
		if v, ok := m[value]; ok && v == 1 {
			return i
		}
	}
	return -1
}
