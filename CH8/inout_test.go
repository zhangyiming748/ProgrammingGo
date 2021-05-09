package CH8

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

/*
8.4输入和输出
 */
func TestIOWrite(t *testing.T) {
	var b bytes.Buffer
	b.Write([]byte("hello "))
	i, err :=fmt.Fprintf(&b,"world!")
	j, erro :=b.WriteTo(os.Stdout)
	t.Logf("%d %d \n",i,j)
	t.Log(err.Error())
	t.Log(erro.Error())
}
/*
简易的curl工具
 */
func TestCurl(t *testing.T) {
	//var r *http.Response
	r,_:=http.Get("https://ss1.baidu.com/-4o3dSag_xI4khGko9WTAnF6hhy/zhidao/pic/item/9a504fc2d5628535c1e6847690ef76c6a7ef6366.jpg")
	file,_:=os.Create("curl.jpg")
	defer file.Close()
	dest:=io.MultiWriter(os.Stdout,file)
	io.Copy(dest,r.Body)
	r.Body.Close()


}
