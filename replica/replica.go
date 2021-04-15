package replica

import "fmt"

type info struct {
	name string
	age  int
}

func master() {
	var i info
	i.name = "zhangyiming"
	i.age = 25
	fmt.Printf("ageP return :%v\n", ageP(i))
	fmt.Printf("name:%v\tage:%v\n", i.name, i.age)
	fmt.Printf("ageP return :%v\n", agePP(&i))
	fmt.Printf("name:%v\tage:%v\n", i.name, i.age)

}
func ageP(i info) info {
	i.age += 5
	return i
}
func agePP(i *info) *info {
	i.age += 5
	return i
}
