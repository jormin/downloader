package helper

import (
	"strings"
	"time"
)

// FormatTime format time from timestamp
func FormatTime(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
}

// FormatDate format date from timestamp
func FormatDate(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("2006-01-02")
}

// RemoveIllegalCharacters remove illegal characters, such as ` `, `/`, `:` etc.
func RemoveIllegalCharacters(str string) string {
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "/", "", -1)
	str = strings.Replace(str, "：", "", -1)
	str = strings.Replace(str, ":", "", -1)
	return str
}
