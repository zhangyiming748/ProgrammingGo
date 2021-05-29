package cow

import ()

func FirstNotRepeatingChar(str string) int {
	s := []byte(str)
	//fmt.Println(s)
	single := []byte{}
	m := make(map[byte]int)
	for _, v := range s {
		if _, ok := m[v]; ok {
			m[v] = m[v] + 1
		} else {
			m[v] = 1
		}
		//fmt.Printf("循环中:%v\n",m)
	}
	for key, value := range m {
		if value == 1 {
			single = append(single, key)
		}
	}
	for idx, val := range s {
		for _, v := range single {
			if val == v {
				return idx

			}
		}
	}
	return -1
}
