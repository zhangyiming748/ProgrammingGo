package Algorithm

import (
	"fmt"
	"strconv"
	"testing"
)

func TestTelegram(t *testing.T) {
	//rand.Seed(time.Now().UnixNano())
	//n := 123
	for i := 0; i < 100000; i++ {
		t.Logf("此次循环数字为%d", i)
		CountNum(i)
		PrintNum(i)
		DescNum(i)
	}

}

func CountNum(n int) (long int) {
	sn := strconv.Itoa(n)
	switch len(sn) {
	case 1:
		fmt.Println("一位数")
		return 1
	case 2:
		fmt.Println("两位数")
		return 2
	case 3:
		fmt.Println("三位数")
		return 3
	case 4:
		fmt.Println("四位数")
		return 4
	case 5:
		fmt.Println("五位数")
		return 5
	default:
		fmt.Println("跳出三界外,不在五行中")
		return 0
	}
}
func PrintNum(n int) {
	sn := strconv.Itoa(n)
	for i, v := range sn {
		fmt.Printf("第%d位数字是%c\n", i+1, v)
	}
}
func DescNum(n int) {
	count := 0
	for n > 0 {
		count += n % 10 //count=3
		n = n / 10      //n=12
		if n < 1 {
			break
		}
		count = count * 10 //count=30
	}
	fmt.Printf("倒序输出的数字是%d\n", count)
}
