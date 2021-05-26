package cow

func MoreThanHalfNum_Solution(numbers []int) int {
	m := make(map[int]int) //map[元素]出现次数
	if len(numbers) == 0 {
		return 0
	}
	if len(numbers) == 1 {
		return numbers[0]
	}
	for _, v := range numbers {
		m[v] += 1
	}
	maxkey := 0
	maxval := 0
	for _, val := range m {
		if val > maxval {
			maxval = val
		}
	}
	for k, value := range m {
		if value == maxval {
			maxkey = k
		}
	}
	return maxkey
}
