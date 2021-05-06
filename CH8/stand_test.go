package CH8

import (
	"log"
	"testing"
)

func TestLog(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile|log.Lshortfile|log.LstdFlags)
	/*
		Ldata 日期
		L
	*/
	log.Println("写到标准日志记录器")
	log.Panicln("Println后panic")
	log.Fatalln("Println后os.exit(1)")

}
