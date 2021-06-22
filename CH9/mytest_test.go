package CH9

import (
	"math/rand"
	"net/http"
	"testing"
	"time"
)

func TestPing(t *testing.T) {
	const (
		checkMark = "\u2713" //✓
		ballotX   = "\u2717" //✗
	)

	url1 := "https://www.baidu.com/"
	//url2:="https://www.google.com/ncr"
	statusCode := 200
	t.Log("test connect")
	{
		//t.Logf("\tWhen checking \"%s\" for status code \"%d\"",url,statusCode)
		{
			resp, err := http.Get(url1)
			if err != nil {
				t.Fatal(ballotX)
			}
			//t.Log(checkMark)
			defer resp.Body.Close()
			if resp.StatusCode == statusCode {
				t.Log(checkMark)

			} else {
				t.Error(ballotX)
			}
		}
	}
}
func BenchmarkRand(b *testing.B) {
	rand.Seed(time.Now().Unix())
	//与测试无关的代码
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := rand.Intn(5000)
		if i > n {
			b.Logf("i = %d bigger tham n = %d\n", i, n)
		} else {
			b.Logf("i = %d smaller tham n = %d\n", i, n)

		}
		time.Sleep(50 * time.Millisecond)
	}
	b.StopTimer()
	//与测试无关的代码
}
