package gStr

import (
	"strconv"
	"lucky/gutil/gLog"
)

//字符串 ->int
func ToInt(str string) (num int) {
	if str == "" {
		gLog.E("strToInt.error", "传入字符串为空")
		return 0
	}
	num, err := strconv.Atoi(str)

	if err != nil {
		gLog.E("strToInt.error", err)
		return 0
	}
	return num
}

//字符串->float32
func ToFloat32(str string) float32 {
	val, err := strconv.ParseFloat(str, 32)

	if err != nil {
		gLog.E("strToFloat.err", err)
		return 0
	}
	return float32(val)
}

//字符串->float64
func ToFloat64(str string) float64 {
	val, err := strconv.ParseFloat(str, 64)

	if err != nil {
		gLog.E("strToFloat64Err", err)
		return 0
	}
	return val
}

func ToInt64(str string) int64 {
	if str == "" {
		gLog.E("strToInt64.error", "str is empty")
		return 0
	}

	retInt, err := strconv.ParseInt(str, 10, 64)

	if err != nil {
		gLog.E("strToInt64.error", "err")
		return 0
	}

	return retInt
}
