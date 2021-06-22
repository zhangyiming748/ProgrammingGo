package strangeSort

import "testing"

/*
描述
现在给出数n和一个1到n的数组，求最少的选择次数，使数组变为升序。

示例1
输入：
4,[4,1,2,3]
复制
返回值：
1
复制
备注：
n<=10^6
数据包含一个整数n和一个含有n个元素的数组，表示从队头到队尾的人的身高。
输出一个整数表示答案。
function wwork( n ,  a ) {
    // write code here
    let min = a[n-1];
    let num = 0;
    for(let i = n - 2; i >= 0; i--){
        let cur = a[i];
        if(a[i] > min){
            num++;
        } else {
            min = a[i];
        }
    }
    return num;
}

*/

func TestWwork(t *testing.T) {
	index := 4
	nums := []int{4, 1, 2, 3}
	ret := wwork(index, nums)
	t.Log(ret)
}
func wwork(n int, a []int) int {
	min := a[n-1]
	num := 0
	for i := n - 2; i >= 0; i-- {
		//cur:=a[i]
		if a[i] > min {
			num++
		} else {
			min = a[i]
		}
	}
	return num
}
