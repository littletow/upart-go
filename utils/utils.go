package utils

import (
	"strconv"
	"time"
)

// Str2Int 字符串转为整数
func Str2Int(str string) int {
	if str == "" {
		return 0
	}
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return i
}

// Str2Int64 字符串转为整数64
func Str2Int64(str string) int64 {
	if str == "" {
		return 0
	}
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

func TS2Str(ts time.Time) string {
	str := ts.Format("2006-01-02 15:04:05")
	return str
}
