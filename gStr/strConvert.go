package gStr

import (
	"strconv"
)

//string -> uint8
func ToUint8(str string) (i uint8) {
	if str == "" {
		return 0
	}

	ui, err := strconv.ParseUint(str, 10, 8)

	if err != nil {
		return 0
	}

	return uint8(ui)
}

//string -> int
func ToInt(str string) (num int) {
	if str == "" {
		return 0
	}
	num, err := strconv.Atoi(str)

	if err != nil {
		return 0
	}
	return num
}

//string -> float32
func ToFloat32(str string) float32 {
	val, err := strconv.ParseFloat(str, 32)

	if err != nil {
		return 0
	}
	return float32(val)
}

//string -> float64
func ToFloat64(str string) float64 {
	val, err := strconv.ParseFloat(str, 64)

	if err != nil {
		return 0
	}
	return val
}

//string -> int64
func ToInt64(str string) int64 {
	if str == "" {
		return 0
	}

	retInt, err := strconv.ParseInt(str, 10, 64)

	if err != nil {
		return 0
	}

	return retInt
}
