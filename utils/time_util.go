package utils

import "time"

//获取当前时间（年月日）
func GetNowTimeShort() string {
	now := time.Now().Format("2006-01-02")
	return now
}

//获取当前时间（年月日时分秒）
func GetNowTimeLong() string {
	now := time.Now().Format("2006-01-02 15:04:05")
	return now
}

//获取当前时间字符串（自定格式）
func GetNowTimeStr() string {
	now := time.Now().Format("20060102150405")
	return now
}