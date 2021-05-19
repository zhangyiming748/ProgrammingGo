package CH8

import (
	"encoding/json"
	"fmt"
	//"strings"
	"testing"
	//"time"
)

func TestMarshal(t *testing.T) {
	//create a map to save
	c := make(map[string]interface{})
	c["name"] = "zen"
	c["title"] = "programmer"
	c["contact"] = map[string]interface{}{
		"home": "415.333.3333",
		"cell": "415.555.5555",
	}
	data, err := json.Marshal(c)
	if err != nil {
		t.Log(err)
	}
	fmt.Println(string(data))
}
func TestMarshalIndent(t *testing.T) {
	//create a map to save
	c := make(map[string]interface{})
	c["name"] = "zen"
	c["title"] = "programmer"
	c["contact"] = map[string]interface{}{
		"home": "415.333.3333",
		"cell": "415.555.5555",
	}
	data, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		t.Log(err)
	}
	fmt.Println(string(data))
}
