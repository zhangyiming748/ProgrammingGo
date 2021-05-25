package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 将给定数组排序
 * @param arr int整型一维数组 待排序的数组
 * @return int整型一维数组
 */
func MySort( arr []int ) []int {

	key := values[0]
	p := 0
	i,j := 0,len(arr)-1
	for i  <= j {

		for j >= p && values[j] >= key {
			j--
		}

		if j >= p {
			values[p] = values[j]
			p = j
		}
		if values[i] <= key && i <= p {
			i ++
		}
		if i < p{
			values[p] = values[i]
			p = i
		}
		values[p] = key
		if p - left > 1{
			quickSort(values, left, p-1)
		}
		if right - p > 1 {
			quickSort(values, p+1, right)
		}
	}

	return arr
}