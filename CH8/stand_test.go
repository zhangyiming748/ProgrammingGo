package CH8

import (
	"encoding/json"
	"fmt"
	"runtime"

	//"encoding/xml"

	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
)

var (
	key = "3e469a45a72b31c120b0baf6609ae877"
	//换成你自己的key
)

type GEO struct {
	Status    string `json:"status"`
	Info      string `json:"info"`
	Infocode  string `json:"infocode"`
	Regeocode struct {
		AddressComponent struct {
			City         string `json:"city"`
			Province     string `json:"province"`
			Adcode       string `json:"adcode"`
			District     string `json:"district"`
			Towncode     string `json:"towncode"`
			StreetNumber struct {
				Number    string `json:"number"`
				Location  string `json:"location"`
				Direction string `json:"direction"`
				Distance  string `json:"distance"`
				Street    string `json:"street"`
			} `json:"streetNumber"`
			Country       string `json:"country"`
			Township      string `json:"township"`
			BusinessAreas []struct {
				Location string `json:"location"`
				Name     string `json:"name"`
				Id       string `json:"id"`
			} `json:"businessAreas"`
			Building struct {
				Name  string `json:"name"`
				Typei string `json:"type"`
			} `json:"building"`
			Neighborhood struct {
				Name  string `json:"name"`
				Typei string `json:"type"`
			} `json:"neighborhood"`
		} `json:"addressComponent"`
		Formatted_address string `json:"formatted_address"`
	} `json:"regeocode"`
}

func TestLog(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile | log.Lshortfile | log.LstdFlags)
	/*
		Ldata 日期
		L
	*/
	log.Println("写到标准日志记录器")
	log.Panicln("Println后panic")
	log.Fatalln("Println后os.exit(1)")

}

func TestLogger(t *testing.T) {
	var (
		Trace   *log.Logger
		Info    *log.Logger
		Warning *log.Logger
		Error   *log.Logger
	)
	file, err := os.OpenFile("error.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("文件打开发生错误", err)
	}
	Trace = log.New(ioutil.Discard, "TRASE: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	Trace.Println("我有一些废话要说")
	Info.Println("我有一些信息要说")
	Warning.Println("我有一些人生的经验要说")
	Error.Println("我要警告你")
}

/*
为了模拟真实的数据构造json
为了解析正确用了一天时间(加上熟悉log包)
*/
func TestEncodeJSON(t *testing.T) {
	var c GEO

	location := "116.0119343,39.66127144"
	JSON := Decode(location)
	err := json.Unmarshal(JSON, &c)

	if err != nil {
		t.Log(err)
	}
	fmt.Printf("status = %s\n", c.Status)
	if c.Regeocode.AddressComponent.City == "" {
		c.Regeocode.AddressComponent.City = "直辖市"
	}
	fmt.Printf("city = %s\n", c.Regeocode.AddressComponent.City)
	fmt.Printf("province = %s\n", c.Regeocode.AddressComponent.Province)
	fmt.Printf("adcode = %s\n", c.Regeocode.AddressComponent.Adcode)
	fmt.Printf("当前位置是%s %s %s %s %s %s\n", c.Regeocode.AddressComponent.Country, c.Regeocode.AddressComponent.Province, c.Regeocode.AddressComponent.District, c.Regeocode.AddressComponent.StreetNumber.Street, c.Regeocode.AddressComponent.StreetNumber.Number, c.Regeocode.AddressComponent.StreetNumber.Direction)

}
func TestTOMap(t *testing.T) {
	var m map[string]interface{}
	location := "116.0119343,39.66127144"
	JSON := Decode(location)
	err := json.Unmarshal(JSON, &m)
	if err != nil {
		t.Log(err)
	}
	fmt.Println(m)

}
func Decode(locat string) []byte {
	url := "https://restapi.amap.com/v3/geocode/regeo?output=json&location=" + locat + "&key=" + key + "&radius=1000&extensions=all"
	response, err := http.Get(url)
	if err != nil {
		//...
	}
	defer response.Body.Close() //在回复后必须关闭回复的主体
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {

	}
	return body
}
func TestIP(t *testing.T) {
	url := "https://restapi.amap.com/v3/ip?output=json&key=" + key
	response, err := http.Get(url)
	if err != nil {
		//...
	}
	defer response.Body.Close() //在回复后必须关闭回复的主体
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Log(err)
	}
	t.Log(string(body))
}

//https://restapi.amap.com/v3/ip?ip=8.8.8.8&output=xml&key=<用户的key>
func TestCPU(t *testing.T) {
	var num int
	//num=runtime.GOMAXPROCS
	t.Logf("cpu core num is : %d", runtime.NumCPU())
	num = 10
	runtime.GOMAXPROCS(10)
	t.Logf("use cpu core num is : %d", num)
}

