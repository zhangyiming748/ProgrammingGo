package CH8

import (
	"encoding/json"
	"fmt"

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
	Status    string    `json:"status"`
	Info      string `json:"info"`
	Infocode  string    `json:"infocode"`
	Regeocode struct {
		AddressComponent struct {
			City         string `json:"city"`
			Province     string `json:"province"`
			Adcode       string    `json:"adcode"`
			District     string `json:"district"`
			Towncode     string `json:"towncode"`
			StreetNumber struct {
				Number    string  `json:"number"`
				Location  string  `json:"location"`
				Direction string  `json:"direction"`
				Distance  string `json:"distance"`
				Street    string  `json:"street"`
			} `json:"street_number"`
			Country       string `json:"country"`
			Township      string `json:"township"`
			BusinessAreas []struct {
				Location string `json:"location"`
				Name     string `json:"name"`
				Id       string    `json:"id"`
			} `json:"business_areas"`
			Building struct {
				Name  string `json:"name"`
				Typei string `json:"type"`
			} `json:"building"`
			Neighborhood struct {
				Name  string `json:"name"`
				Typei string `json:"type"`
			} `json:"neighborhood"`
		} `json:"address_component"`
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

func TestEncodeJSON(t *testing.T) {
	var c GEO
	location := "116.36800384499999,39.91048431388889"
	JSON := Decode(location)
	err := json.Unmarshal(JSON, &c)
	if err != nil {
		t.Log(err)
	}
	fmt.Println(c.Regeocode.Formatted_address)
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