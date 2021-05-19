package CH9

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"testing"
)

func TestParseJSON(t *testing.T) {
	filename := "static.json"
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	json := string(data)
	value := gjson.Get(json, "retData.list")
	println(value.String())

}
