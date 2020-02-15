package utils

import (
	"strconv"
	"time"
)

func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func IntToString(a int) string {
	return strconv.Itoa(a)
}

func CurrentTimestamp() int64 {
	t := time.Now()
	return t.Unix()
}

func TimeStringToTimestamp(dateTimeStr string) int64 {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	tt, _ := time.ParseInLocation("2006-01-02 15:04:05", dateTimeStr, loc)
	return tt.Unix()
}
