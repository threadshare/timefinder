package timefinder

import (
	"log"
	"strconv"
)

// stringToInt 字符串转INT
func stringToInt(str string) int {
	res, err := strconv.Atoi(str)
	if err != nil {
		log.Printf("Error converting string to int: %v", err)
		return 0 // 或者返回一个错误值，取决于你的需求
	}
	return res
}

// includes 检查切片中是否包含指定值
func includes(list []string, value string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}
