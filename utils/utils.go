package utils

import (
	"crypto/md5"
	"encoding/hex"
	"net"
	"strconv"
	"time"
)

const WsShareKey string = "20!I@LOVE$CHINA#24"

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

// EncodeMD5 md5编码
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}

// 获取本机MAC地址
func GetMacAddr() string {
	ints, err := net.Interfaces()
	if err != nil {
		return ""
	}

	i0 := ints[0]
	mac := i0.HardwareAddr.String()
	return mac
}

// 生成vcode
func GenVcode(mac string) string {
	str := WsShareKey + mac
	mstr := EncodeMD5(str)
	return mstr
}
