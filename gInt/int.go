package gInt

//int to bool
func ToBool(i int) bool {
	if i == 0 {
		return false
	} else {
		return true
	}
}

//int abs
func IntAbs(i int) int {
	y := i >> 31
	return (i ^ y) - y
}
