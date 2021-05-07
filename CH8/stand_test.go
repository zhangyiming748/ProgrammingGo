package CH8

import (
	"encoding/json"

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
	status    string `json:"status"`
	regeocode string `json:"regeocode"`
	info      string `json:"info"`
	infocode  int    `json:"infocode"`
}
type regeocode struct {
	roads             []road      `json:"roads"`
	roadinters        []roadinter `json:"roadinters"`
	formatted_address string      `json:"formatted_address"`
	addressComponent  string      `json:"address_component"`
	aois              []aoi       `json:"aois"`
	pois              []poi       `json:"pois"`
}
type road struct {
	id        string `json:"id"`
	location  string `json:"location"`
	direction string `json:"direction"`
	name      string `json:"name"`
	distance  string `json:"distance"`
}
type roadinter struct {
	second_name string `json:"second_name"`
	first_id    string `json:"first_id"`
	second_id   string `json:"second_id"`
	Location    string `json:"location"`
	distance    string `json:"distance"`
	first_name  string `json:"first_name"`
	direction   string `json:"direction"`
}
type aoi struct {
	area     string `json:"area"`
	typei    string `json:"type"`
	id       string `json:"id"`
	location string `json:"location"`
	adcode   string `json:"adcode"`
	name     string `json:"name"`
	distance string `json:"distance"`
}
type poi struct {
	id           string `json:"id"`
	direction    string `json:"direction"`
	businessarea string `json:"businessarea"`
	address      string `json:"address"`
	poiweight    string `json:"poiweight"`
	name         string `json:"name"`
	Location     string `json:"location"`
	distance     string `json:"distance"`
	tel          string `json:"tel"`
	typei        string `json:"typei"`
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
	JSON:=Decode(location)
	err:=json.Unmarshal(JSON,&c)
	if err !=nil{
		t.Log(err)
	}
	t.Logf("decode之后:%v",JSON)
}
func Decode(locat string)[]byte {
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