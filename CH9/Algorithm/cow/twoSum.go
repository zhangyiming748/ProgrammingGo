package cow


func twoSum(numbers []int, target int) []int {
	if len(numbers) <= 2 {
		return nil
	}
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > target {
			continue
		}
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				return []int{i + 1, j + 1}
			}
		}
	}
	return nil
}
