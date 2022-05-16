package util

import (
	"os"
	"strings"
)

// StrFirstToUpper 首字母转换成大写
func StrFirstToUpper(str string) string {
	if len(str) < 1 {
		return ""
	}

	strArray := []rune(str)

	if strArray[0] >= 97 && strArray[0] <= 122 {
		strArray[0] -= 32
	}

	return string(strArray)
}

// StrFirstToLower 首字母转换成小写
func StrFirstToLower(str string) string {
	if len(str) < 1 {
		return ""
	}

	strArray := []rune(str)

	if strArray[0] >= 65 && strArray[0] <= 90 {
		strArray[0] += 32
	}

	return string(strArray)
}

// StrCamel 将下划线连接的字串改为驼峰
func StrCamel(str string) string {
	tmp := strings.Split(str, "_")
	var result string
	for _, v := range tmp {
		result += StrFirstToUpper(v)
	}

	return result
}

func PathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}
