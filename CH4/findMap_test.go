package CH4

import "testing"

func TestFindMap(t *testing.T) {
	m := make(map[string]int)
	m["zero"] = 0
	if v, ok := m["zero"]; ok {
		t.Logf("zero存在,v = %v\tok = %v\nok的类型是%T", v, ok, ok)
	}
	if key, val := m["zero"]; val {
		t.Logf("zero存在,v = %v\tok = %v\nval的类型是%T", key, val, val)
	}
	for k, value := range m {
		t.Logf("zero存在,k = %v\tvalue = %v\nval的类型是%T", k, value, value)
	}

}
