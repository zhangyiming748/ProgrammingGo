package CH9

import (
	"fmt"
	"strings"
	"testing"
	"time"
)



func TestCharacter(t *testing.T) {
	t.Logf("")
	for i:=0;i<16;i++{
		for j:=0;j<16;j++{
			for p:=0;p<16;p++{
				for q:=0;q<16;q++{
					fi:=convert(i)
					//t.Logf("fi = %s",fi)
					se:=convert(j)
					th:=convert(p)
					fo:=convert(q)
					t.Logf("%s",strings.Join([]string{fi,se,th,fo},""))
					ret:=fmt.Sprintf("\"\\u%s\"",strings.Join([]string{fi,se,th,fo},""))
					fmt.Printf("ret = %s",ret)
					fmt.Println(ret)
					time.Sleep(time.Second)
				}
			}
		}
	}
}
func convert(i int)string{
	switch i {
	case 0:
		return "0"
	case 1:
		return "1"
	case 2:
		return "2"
	case 3:
		return "3"
	case 4:
		return "4"
	case 5:
		return "5"
	case 6:
		return "6"
	case 7:
		return "7"
	case 8:
		return "8"
	case 9:
		return "9"
	case 10:
		return "A"
	case 11:
		return "B"
	case 12:
		return "C"
	case 13:
		return "D"
	case 14:
		return "E"
	case 15:
		return "F"
	default:
		return ""
	}

}
func TestConvert(t *testing.T) {
	//i:=14
	//r:=convert(i)
	//r := "A"+"C"+"C"+"D"
	//ret:=fmt.Sprintf("%x",r)
	//ret:=strcon
	//t.Log(ret)
}