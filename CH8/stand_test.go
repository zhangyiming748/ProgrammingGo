package CH8

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

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
