package main

func ReString(str string) string {
	data := []byte{}
	newdata := []byte{}
	data = []byte(str)
	for len(data) != 0 {
		newdata = append(newdata, data[len(data)-1])
		data = data[:len(data)-1]
	}
	return string(newdata)
}
