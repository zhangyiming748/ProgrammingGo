package CH8

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

const ()

func TestWeather(t *testing.T) {
	m := make(map[string]string)
	m["北京市石景山区"] = "110107"
	m["黑龙江省大庆市让胡路区"] = "230604"
	Weather(m["北京市石景山区"])

	//fmt.Println("\u00B0")
}
func Weather(adcode string) {
	var m map[string]interface{}
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
	err2 := json.Unmarshal(body, &m)
	if err2 != nil {
		log.Println(string(body))
	}
	//log.Printf("开始\n%v\n结束", string(body))
	//t.Logf("lives : %s",m["lives"])
	//fmt.Println(m["lives"].([]interface{})[0].(map[string]interface{})["city"])
	//city := m["forecasts"].([]interface{})[0].(map[string]interface{})["city"]
	//province := m["forecasts"].([]interface{})[0].(map[string]interface{})["province"]
	reporttime := m["forecasts"].([]interface{})[0].(map[string]interface{})["reporttime"]
	fmt.Printf("%v发布\n", reporttime)
	today_date := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[0].(map[string]interface{})["date"]
	today_week := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[0].(map[string]interface{})["week"]

	switch today_week {
	case "1":
		today_week = "一"
	case "2":
		today_week = "二"
	case "3":
		today_week = "三"
	case "4":
		today_week = "四"
	case "5":
		today_week = "五"
	case "6":
		today_week = "六"
	case "7":
		today_week = "日"
	}
	today_dayweather := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[0].(map[string]interface{})["dayweather"]
	today_nightweather := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[0].(map[string]interface{})["nightweather"]
	today_daytemp := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[0].(map[string]interface{})["daytemp"]
	today_nighttemp := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[0].(map[string]interface{})["nighttemp"]
	today_daywind := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[0].(map[string]interface{})["daywind"]
	today_nightwind := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[0].(map[string]interface{})["nightwind"]
	today_daypower := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[0].(map[string]interface{})["daypower"]
	today_nightpower := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[0].(map[string]interface{})["nightpower"]
	fmt.Printf("%v 星期%v\t白天%v最高温度%v\u00B0C %v风%v级\t夜间%v最低温度%v\u00B0C%v风%v级\n", today_date, today_week, today_dayweather, today_daytemp, today_daywind, today_daypower, today_nighttemp, today_nightweather, today_nightwind, today_nightpower)

	after_24_date := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[1].(map[string]interface{})["date"]
	after_24_week := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[1].(map[string]interface{})["week"]

	switch after_24_week {
	case "1":
		after_24_week = "一"
	case "2":
		after_24_week = "二"
	case "3":
		after_24_week = "三"
	case "4":
		after_24_week = "四"
	case "5":
		after_24_week = "五"
	case "6":
		after_24_week = "六"
	case "7":
		after_24_week = "日"
	}

	after_24_dayweather := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[1].(map[string]interface{})["dayweather"]
	after_24_nightweather := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[1].(map[string]interface{})["nightweather"]
	after_24_daytemp := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[1].(map[string]interface{})["daytemp"]
	after_24_nighttemp := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[1].(map[string]interface{})["nighttemp"]
	after_24_daywind := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[1].(map[string]interface{})["daywind"]
	after_24_nightwind := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[1].(map[string]interface{})["nightwind"]
	after_24_daypower := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[1].(map[string]interface{})["daypower"]
	after_24_nightpower := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[1].(map[string]interface{})["nightpower"]

	fmt.Printf("%v 星期%v\t白天%v最高温度%v\u00B0C %v风%v级\t夜间%v最低温度%v\u00B0C%v风%v级\n", after_24_date, after_24_week, after_24_dayweather, after_24_daytemp, after_24_daywind, after_24_daypower, after_24_nighttemp, after_24_nightweather, after_24_nightwind, after_24_nightpower)

	after_48_date := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[2].(map[string]interface{})["date"]
	after_48_week := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[2].(map[string]interface{})["week"]

	switch after_48_week {
	case "1":
		after_48_week = "一"
	case "2":
		after_48_week = "二"
	case "3":
		after_48_week = "三"
	case "4":
		after_48_week = "四"
	case "5":
		after_48_week = "五"
	case "6":
		after_48_week = "六"
	case "7":
		after_48_week = "日"
	}
	after_48_dayweather := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[2].(map[string]interface{})["dayweather"]
	after_48_nightweather := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[2].(map[string]interface{})["nightweather"]
	after_48_daytemp := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[2].(map[string]interface{})["daytemp"]
	after_48_nighttemp := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[2].(map[string]interface{})["nighttemp"]
	after_48_daywind := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[2].(map[string]interface{})["daywind"]
	after_48_nightwind := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[2].(map[string]interface{})["nightwind"]
	after_48_daypower := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[2].(map[string]interface{})["daypower"]
	after_48_nightpower := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[2].(map[string]interface{})["nightpower"]

	fmt.Printf("%v 星期%v\t白天%v最高温度%v\u00B0C %v风%v级\t夜间%v最低温度%v\u00B0C%v风%v级\n", after_48_date, after_48_week, after_48_dayweather, after_48_daytemp, after_48_daywind, after_48_daypower, after_48_nighttemp, after_48_nightweather, after_48_nightwind, after_48_nightpower)

	//after_48
	after_72_date := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[3].(map[string]interface{})["date"]
	after_72_week := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[3].(map[string]interface{})["week"]

	switch after_72_week {
	case "1":
		after_72_week = "一"
	case "2":
		after_72_week = "二"
	case "3":
		after_72_week = "三"
	case "4":
		after_72_week = "四"
	case "5":
		after_72_week = "五"
	case "6":
		after_72_week = "六"
	case "7":
		after_72_week = "日"
	}
	after_72_dayweather := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[3].(map[string]interface{})["dayweather"]
	after_72_nightweather := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[3].(map[string]interface{})["nightweather"]
	after_72_daytemp := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[3].(map[string]interface{})["daytemp"]
	after_72_nighttemp := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[3].(map[string]interface{})["nighttemp"]
	after_72_daywind := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[3].(map[string]interface{})["daywind"]
	after_72_nightwind := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[3].(map[string]interface{})["nightwind"]
	after_72_daypower := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[3].(map[string]interface{})["daypower"]
	after_72_nightpower := m["forecasts"].([]interface{})[0].(map[string]interface{})["casts"].([]interface{})[3].(map[string]interface{})["nightpower"]

	fmt.Printf("%v 星期%v\t白天%v最高温度%v\u00B0C %v风%v级\t夜间%v最低温度%v\u00B0C%v风%v级\n", after_72_date, after_72_week, after_72_dayweather, after_72_daytemp, after_72_daywind, after_72_daypower, after_72_nighttemp, after_72_nightweather, after_72_nightwind, after_72_nightpower)

	//24,48,
}
