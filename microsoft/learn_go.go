package microsoft
/*
https://docs.microsoft.com/zh-cn/learn/modules/go-get-started
 */

import (
	"fmt"
	"math"

)

func hello() {
	fmt.Println("Hello World!")
}
func typeInt() {
	var integer8 int8 = 127
	var integer16 int16 = 32767
	var integer32 int32 = 2147483647
	var integer64 int64 = 9223372036854775807
	println(integer8, integer16, integer32, integer64)
	myRune := 'G'
	println(myRune)
}
func typeFloat() {
	var f32 float32 = 2147483647
	var f64 float64 = 9223372036854775807
	println(f32, f64)
	println(math.MaxFloat32, math.MaxFloat64)

}
func typeString() {
	var firstName string = "John"
	lastName := "Doe"
	println(firstName, lastName)
}

func calc(i int) {
	total := internalSum(i)
	println(total)
}
func internalSum(i int)int{
	var sum int
	for ;i>0;i--{
		sum=sum+i
	}
	return sum
}