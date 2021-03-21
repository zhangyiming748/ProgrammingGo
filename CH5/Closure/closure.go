package Closure

import (
	"fmt"
	"strings"
)

func closure()  {
		addPng:= func(name string)string {
			return name+".png"
		}
		addJpg:= func(name string)string {
			return name+".jpg"
		}
		fmt.Println(addPng("filename"))
		fmt.Println(addJpg("filename"))
}
func factoryfunc(suffix string)func(name string)string {
	return func(name string) string {
		if !strings.HasSuffix(name,suffix){
			return name+suffix
		}
		return name
		
	}
}