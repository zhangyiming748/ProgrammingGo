package main

import (

	"io/ioutil"
	"log"

	"net/http"
)


var (
	key = "3e469a45a72b31c120b0baf6609ae877"
	//换成你自己的key
)
func main() {

	Weather("100000")

	//fmt.Println("\u00B0")
}
func Weather(adcode string) {
	//var m map[string]interface{}
	url := "https://restapi.amap.com/v3/weather/weatherInfo?output=json&extensions=all&city=" + adcode + "&key=" + key
	response, err := http.Get(url)
	if err != nil {
		//...
	}
	defer response.Body.Close() //在回复后必须关闭回复的主体
	body, err1 := ioutil.ReadAll(response.Body)
	if err1 != nil {
		log.Println(err1)
	}
	//err2 := json.Unmarshal(body, &m)
	//if err2 != nil {
	//	log.Println(string(body))
	//}
	log.Printf("开始\n%v\n结束", string(body))

}
