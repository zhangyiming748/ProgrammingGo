package CH9

import (
	"net/http"
	"testing"
)

func TestPing(t *testing.T) {
	const (
		checkMark = "\u2713" //✓
		ballotX = "\u2717" //✗
	)


	url1:="https://www.baidu.com/"
	//url2:="https://www.google.com/ncr"
	statusCode:=200
	t.Log("test connect")
	{
		//t.Logf("\tWhen checking \"%s\" for status code \"%d\"",url,statusCode)
		{
			resp ,err:=http.Get(url1)
			if err !=nil{
				t.Fatal(ballotX)
			}
			//t.Log(checkMark)
			defer resp.Body.Close()
			if resp.StatusCode==statusCode{
				t.Log(checkMark)

			}else {
				t.Error(ballotX)
			}
		}
	}
}
