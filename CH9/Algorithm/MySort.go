package Algorithm

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}
func quickSort(values []int, left, right int) {
	key := values[left] //取出第一项
	p := left
	i, j := left, right
	for i <= j {
		//由后开始向前搜索(j--)，找到第一个小于key的值values[j]
		for j >= p && values[j] >= key {
			j--
		}
		//第一个小于key的值 赋给 values[p]
		if j >= p {
			values[p] = values[j]
			p = j
		}
		if values[i] <= key && i <= p {
			i++
		}
		if i < p {
			values[p] = values[i]
			p = i
		}
		values[p] = key
		if p-left > 1 {
			quickSort(values, left, p-1)
		}
		if right-p > 1 {
			quickSort(values, p+1, right)
		}
	}
}
