package Exercise

import (

	"reflect"
	"testing"
)

func TestUniqueInts(t *testing.T) {
	nums:=[]int{9,1,9,5,4,4,2,1,5,4,8,8,4,3,6,9,5,7,5}
	expect :=[]int{9,1,5,4,2,8,3,6,7}
	ret:=uniqueInts(nums)
	if reflect.DeepEqual(expect,ret){
		t.Log("PASS")
	}else {
		t.Logf("expect is %v\tbut ret is %v\n",expect,ret)
	}
}
func TestFlatten(t *testing.T) {
	nums := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}, {17, 18, 19, 20}}
	ret := flatten(nums)
	t.Log(ret)
}
func TestMake2D(t *testing.T) {
	nums:=[]int{9,1,9,5,4,4,2,1,5,4,8,8,4,3,6,9,5,7,5}
	ret :=make2d(nums,3)
	t.Log(ret)

}