package GCC

import (
	"fmt"
	"runtime"
	"time"
)

func gcc() {
	for i := 0; i < 60; i++ {
		t := time.Now().Nanosecond()
		fmt.Printf(time.Now().Format("2006-01-02 15:04:05\n"))
		fmt.Println(time.Duration(t))
		if i%10 == 0 {
			runtime.GC()
		}
	}
}
