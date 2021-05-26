package cow

type node struct {
	index int
	left  int
	right int
}

func mount(nums []int) int {
	index := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			index = i
		}
	}
	return index + 1
}
