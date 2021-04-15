package Exercise

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

var (
	iniData = []string{
		"; Cut down copy of Mozilla application.ini file",
		"",
		"[App]",
		"Vendor=Mozilla",
		"Name=Iceweasel",
		"Profile=mozilla/firefox",
		"Version=3.5.16",
		"[Gecko]",
		"MinVersion=1.9.1",
		"MaxVersion=1.9.1.*",
		"[XRE]",
		"EnableProfileMigrator=0",
		"EnableExtensionManager=1",
	}
)

func ParseIni(iniData []string) map[string]map[string]string {
	ini := make(map[string]map[string]string)
	group := "General"
	for _, data := range iniData {
		data := strings.TrimSpace(data)
		if data == "" || strings.HasPrefix(data, ";") {
			continue
		}

		if strings.HasPrefix(data, "[") && strings.HasSuffix(data, "]") {
			group = data[1 : len(data)-1]
		} else if i := strings.Index(data, "="); i > -1 {
			key := data[:i]
			value := data[i+len("="):]
			if _, found := ini[group]; !found {
				ini[group] = make(map[string]string)
			}
			ini[group][key] = value
		} else {
			log.Println("some thing error")
		}
	}
	return ini
}
func PrintIni(ini map[string]map[string]string) {
	groups := make([]string, 0, len(ini))
	for group := range ini {
		groups = append(groups, group)
	}
	sort.Strings(groups)
	for i, group := range groups {
		fmt.Printf("[%s]\n", group)
		keys := make([]string, 0, len(ini[group]))
		for key := range ini[group] {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			fmt.Printf("%s=%s\n", key, ini[group][key])
		}
		if i+1 < len(groups) {
			fmt.Println()
		}
	}
}

func uniqueInts(s []int) []int {
	newSlice := make([]int, 0)
	exist := make(map[int]bool)
	for _, v := range s {
		if _, ok := exist[v]; !ok {
			exist[v] = true
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

func flatten(ss [][]int) []int {
	var ret []int
	for _, s := range ss {
		for _, v := range s {
			ret = append(ret, v)
		}
	}
	return ret
}
func make2d(s []int, n int) [][]int {
	var ss [][]int
	//group:=len(s)/n //多少组

	for i := 0; i < len(s); i++ {
		inner := []int{}
		for j := 0; j < n; j++ {
			inner = append(inner, s[i])
		}
		ss = append(ss, inner)
	}

	return ss
}
