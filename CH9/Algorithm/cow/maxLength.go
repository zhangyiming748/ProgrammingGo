package cow

import (
	//l "ProgrammingGo/log"

	"ProgrammingGo/log"
)

func maxLength(arr []int) int {
	max := 0
	for i := 0; i < len(arr); i++ {
		tmpmax := 0
		for j := i + 1; j < len(arr)-1; j++ {
			log.Debug.Println(i, j)
			if arr[i] == arr[j] {
				tmpmax = j - i + 1
				if tmpmax > max {
					max = tmpmax
					break
				}
			}
			max = len(arr) - i

		}

	}
	return max
}
func resetMap(m map[int]bool) {
	for k, _ := range m {
		m[k] = false
	}
}
