package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)
func init(){
	log.SetFlags(2 | 16)
}

func main() {
	var (
		start string
		end   string
	)
	log.Println("input start time")
	if _, err := fmt.Scanln(&start); err != nil {
		panic("输入格式有误(HH:MM:SS)")
	}
	log.Println("input end time")
	if _, err := fmt.Scanln(&end); err != nil {
		panic("输入格式有误(HH:MM:SS)")
	}

	//time1 := "2015-03-20 08:50:29"
	//time2 := "2015-03-20 09:04:25"
	tstart := strings.Join([]string{"2015-03-20", start}, " ")
	tend := strings.Join([]string{"2015-03-20", end}, " ")
	//先把时间字符串格式化成相同的时间类型
	t1, err := time.Parse("2006-01-02 15:04:05", tstart)
	t2, err := time.Parse("2006-01-02 15:04:05", tend)
	if err == nil && t1.Before(t2) {
		//处理逻辑
		log.Println("true")
	}
	t3 := t2.Sub(t1)
	t := t3.String()
	log.Println(t)
	log.Printf("%T", t)
	res := parseTime(t)
	log.Printf("%s", res)
}
func parseTime(s string) string {
	sb := []byte(s)
	resb := []byte{}
	if strings.Contains(s, "h") {
		for _, v := range sb {
			if v == 'h' || v == 'm' || v == 's' {
				continue
			} else {
				resb = append(resb, v)
			}
		}

	} else if strings.Contains(s, "m") {
		for _, v := range sb {
			if v == 'm' || v == 's' {
				continue
			} else {
				resb = append(resb, v)
			}
		}
	} else {
		for _, v := range sb {
			if v == 's' {
				continue
			} else {
				resb = append(resb, v)
			}
		}
	}
	log.Printf("resb=%s", resb)
	return string(resb)
}
