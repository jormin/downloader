package helper

import (
	"regexp"
	"time"
)

// 从Url中获取Code
func GetCodeFromUrl(url string, index int) string {
	return regexp.MustCompile("[0-9]+").FindAllString(url, -1)[index]
}

// 格式化时间戳
func FormatTime(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
}
