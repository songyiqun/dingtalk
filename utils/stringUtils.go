package utils

import "strconv"

func IntToString(v int64) string {
	return strconv.FormatInt(v, 10)
}

func StrToInt64(v string) int64 {
	result, _ := strconv.ParseInt(v, 10, 64)
	return result
}
