package gFloat

import (
	"fmt"
	"github.com/lucky-lee/gutil/gStr"
	"strconv"
)

//float64 save point after num
func Float64SavePointAfter(f float64, num int) float64 {
	str := fmt.Sprintf("%."+strconv.Itoa(num)+"f", f)
	return gStr.ToFloat64(str)
}

//float32 save point after num
func Float32SavePointAfter(f float32, num int) float32 {
	str := fmt.Sprintf("%."+strconv.Itoa(num)+"f", f)
	return gStr.ToFloat32(str)
}

//float64 save point after two
func Float64SaveAfterTwo(f float64) float64 {
	return Float64SavePointAfter(f, 2)
}

//float32 save point after two
func Float32SaveAfterTwo(f float32) float32 {
	return Float32SavePointAfter(f, 2)
}
