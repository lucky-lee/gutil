package gInt

import "fmt"

//int to bool
func ToBool(i int) bool {
	if i == 0 {
		return false
	}

	return true
}

//int abs
func IntAbs(i int) int {
	y := i >> 31
	return (i ^ y) - y
}

//uint8 to bool
func Uint8ToBool(i uint8) bool {
	if i == 0 {
		return false
	}

	return true
}

//int类型切片 转string
//ints int类型切片
//symbol 符号
func IntsToStr(ints []int, symbol string) (str string) {
	len := len(ints)

	if len == 0 {
		return
	}

	for k, v := range ints {
		if k == len-1 {
			str += fmt.Sprintf("%d", v)
		} else {
			str += fmt.Sprintf("%d%s", v, symbol)
		}
	}

	return
}
