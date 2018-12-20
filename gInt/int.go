package gInt

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
