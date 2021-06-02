package cow

import (
	. "ProgrammingGo/log"
	"log"

	"math"
	//"time"
)

func longestCommonPrefix(strs []string) string {
	same := 0
	//i := 1
	//j := 0
	//s := moreThen(strs)
	//for i < len(strs) || j <= s {
	//	if strs[i][j] != strs[i-1][j] {
	//		break
	//	} else {
	//		same++
	//	}
	//}
	//l:=len(strs)
	//same:=0
	//sb, min := moreThen(strs)
	//for i := 0; i < min; i++ {//0,1,2
	//	for j := 0; j <l; j++ {//0,1,2,3,4
	//		if sb[i]!=byte(strs[j][i]){
	//			Debug.Println("final",same)
	//		}
	//		same++
	//	}
	//}
	//bb := []byte{}
	//for i,v:=range strs{
	//	bv:=[]byte(v)
	//	bb=append(bb,v)
	//}

	//same:=0
	//s, min := moreThen(strs)
	//for i := 0; i < min; i++ {
	//	for j := 0; j < len(strs); j++ {
	//		if s[i] == strs[j][i] {
	//			same++
	//			Debug.Printf("i=%v\tj=%v\ts[i]=%v\tstrs[j][i]=%v\n", i, j, s[i], strs[j][i])
	//		}
	//	}
	//
	//}
	//Debug.Println(same)

	log.Println(same)
	return ""
}
//func moreThen(strs []string) int {
//	min := math.MaxInt32
//
//	for _, v := range strs {
//		if len(v) < min {
//			min = len(v)
//
//			Debug.Printf("min = %v\ts = %v\n", min, v)
//		}
//	}
//
//	return min
//}
