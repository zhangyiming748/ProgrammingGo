package cow

import (
	"fmt"
	"sort"
	"testing"
)

/*
给你一个含有n个元素的数组arr[i]，问这个数组的中位数大还是平均数大，如果中位数更大输出1，如果平均数更大输出-1，如果中位数和平均数相等输出0

示例1
输入：
[1,3,4]
复制
返回值：
1
复制
说明：
中位数3，平均数约等于2.67，所以输出1
示例2
输入：
[7,4,8,11]
复制
返回值：
0
复制
说明：
中位数7.5，平均数7.5，所以输出0
示例3
输入：
[6,6,6,6,5,8]
复制
返回值：
-1
复制
说明：
中位数6，平均数约等于6.17，所以输出-1
*/

func Answerofjudge(arr []int) int {
	i := median(arr)
	j := average(arr)
	if i > j {
		return 1
	} else if i < j {
		return -1
	} else {
		return 0
	}
}
func median(nums []int) float64 {
	sort.Ints(nums)
	avg := 0
	l := len(nums)
	var med float64
	if l%2 == 0 {
		a := float64(nums[l/2])

		b := float64(nums[l/2-1])
		fmt.Printf("a=%v\tb=%v\n", a, b)
		med = (a + b) / 2

	} else {
		avg = nums[l/2]
		fmt.Println(avg)
		med = float64(avg)
	}
	fmt.Printf("%v的中位数是%v\n", nums, avg)
	return med
}
func average(nums []int) float64 {
	sum := 0.0
	l := float64(len(nums))
	for _, v := range nums {
		sum += float64(v)
	}
	var avg float64
	avg = sum / l
	fmt.Printf("%v的平均数是%v\n", nums, avg)
	return avg
}
func TestAnswerofjudge(t *testing.T) {
	t1 := []int{1, 3, 4}
	t2 := []int{7, 4, 8, 11}
	t3 := []int{6, 6, 6, 6, 5, 8}
	ret1 := Answerofjudge(t1)
	ret2 := Answerofjudge(t2)
	ret3 := Answerofjudge(t3)
	t.Logf("t1=%v\tt2=%v\tt3=%v\n", ret1, ret2, ret3)
}
