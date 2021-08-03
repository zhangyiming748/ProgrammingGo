package cow

import "fmt"

func foundOnceNumber(arr []int, k int) int {
	m := make(map[int]int) //key是数组的值value是出现次数
	for _, v := range arr {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	for key, val := range m {
		if val == 1 {
			return key
		}
	}
	return 0
}
func main() {
	arr := []int{5, 4, 1, 1, 5, 1, 5}
	k := 3
	ret := foundOnceNumber(arr, k)
	fmt.Println(ret)
}
