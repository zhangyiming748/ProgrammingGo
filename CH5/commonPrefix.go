package CH5

import (
	"bytes"
	"path/filepath"
	"strings"
)

//接收字符串返回共同前缀
func CommonPrefix(texts []string) string {
	components := make([][]rune, len(texts))
	//创建二维int32
	for i, text := range texts {
		components[i] = []rune(text)
	}
	if len(components) == 0 || len(components[0]) == 0 {
		return ""
	}
	var common bytes.Buffer
FINISH:
	for column := 0; column < len(components[0]); column++ {
		char := components[0][column]
		for row := 1; row < len(components); row++ {
			if column >= len(components[row]) ||
				components[row][column] != char {
				break FINISH
			}
		}
		common.WriteRune(char)
	}
	return common.String()
}
func CommonPathPrefix(paths []string) string {
	const separator = string(filepath.Separator)
	components := make([][]string, len(paths))
	for i, path := range paths {
		components[i] = strings.Split(path, separator)
		if strings.HasPrefix(path, separator) {
			components[i] = append([]string{separator}, components[i]...)
		}
	}
	if len(components) == 0 || len(components[0]) == 0 {
		return ""
	}
	var common []string
FINISH:
	for column := range components[0] {
		part := components[0][column]
		for row := 1; row < len(components); row++ {
			if len(components[row]) == 0 ||
				column >= len(components[row]) ||
				components[row][column] != part {
				break FINISH
			}
		}
		common = append(common, part)
	}
	return filepath.Join(common...)
}
