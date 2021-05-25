package cow

import (
	"fmt"

	"math"
)

func GetLeastNumbers_Solution(input []int, k int) []int {
	nums:=&input
	result:=[]int{}
	for i := 0; i<k; i++ {
		newline,ret:=delMin(*nums)
		nums=&newline
		result=append(result,ret)

	}
	return result
}

type Min struct {
	index int
	value int
}

func findMin(nums []int) Min {
	m := make(map[int]int)
	var ans Min
	ans.value = math.MaxInt32
	for i, v := range nums {
		m[i] = v
	}

	for key, val := range m {
		if val < ans.value {
			ans.value = val
			ans.index = key
		}
	}
	return ans
}
func delMax(nums []int)[]int{
	var maxi int
	var maxv int
	newLine:=[]int{}
	for i,v:=range nums{
		if v>maxi{
			maxv=v
			maxi=i
		}
	}
	newLine=append(nums[:maxi],nums[maxi+1:]...)
	fmt.Println("del max :",maxv)
	return newLine
}
func delMin(s []int) ([]int, int) {
	mini := 0
	minv := math.MaxInt32
	for i, v := range s {
		if v < minv {
			minv = v
			mini = i
			fmt.Printf("此次循环minv=%d,mini=%d", minv, mini)
		}
	}
	newer := append(s[:mini], s[mini+1:]...)
	fmt.Println(newer)
	return newer, minv
}