package cow

import (
	"fmt"
	"strconv"
	"strings"
)

func EffectiveIp(ip string) string {
	if strings.Contains(ip, ".") {
		if isIPV4(ip) {
			return "IPv4"
		}

	}
	if strings.Contains(ip, ":") {
		if isIPV6(ip) {
			return "IPv6"
		}
	}
	return "Neither"
}

//"172.16.254.1"
func isIPV4(ip string) bool {
	s := strings.Split(ip, ".")
	if len(s) < 4 { //不足四组数字
		return false
	}
	for i, v := range s {
		val, _ := strconv.Atoi(v)
		if val < 0 && val > 255 {
			return false
		}
		fmt.Printf("第%d组数字%s符合要求\n", i, v)
	}
	return true
}
func isIPV6(ip string) bool {
	if strings.Contains(ip, "::") {
		return false
	}
	s := strings.Split(ip, ":")
	for _, v := range s {
		if !legal(v) {
			return false
		}
	}
	return true
}
func legal(s string) bool {
	b := []byte(s)
	for i, v := range b {
		if v >= 48 && v <= 57 { //0-9
			return true
		}
		if v >= 65 && v <= 71 { //A-F
			return true
		}
		if v >= 97 && v <= 103 { //a-f
			return true
		}
		fmt.Printf("第%d个字符不合法", i)
	}
	return false
}
