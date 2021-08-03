package singleNumber

import "testing"
/*
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
 */
func singleNumber1(nums []int)int{
	m := make(map[int]int) //value=出现次数
	for _, v := range nums {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	for key, value := range m {
		if value == 1 {
			return key
		}
	}
	return 0
}
/*
给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。
*/
func singleNumber2(nums []int) int {
	m := make(map[int]int) //value=出现次数
	for _, v := range nums {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	for key, value := range m {
		if value == 1 {
			return key
		}
	}
	return 0
}

/*
给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。你可以按 任意顺序 返回答案。
*/
func singleNumber3(nums []int) []int {
	m := make(map[int]int) //value=出现次数
	for _, v := range nums {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	ans := []int{}
	for key, value := range m {
		if value == 1 {
			ans = append(ans, key)
		}
	}
	return ans
}
func TestUnit(t *testing.T) {
	nums1 := []int{2, 2, 3, 2}           //3
	nums2 := []int{0, 1, 0, 1, 0, 1, 99} //99
}
