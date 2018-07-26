package gFloat

import (
	"strconv"
	"fmt"
	"github.com/lucky-lee/gutil/gStr"
)

//float64保留小数后几位
func Float64SavePointAfter(f float64, num int) float64 {
	str := fmt.Sprintf("%."+strconv.Itoa(num)+"f", f)
	return gStr.ToFloat64(str)
}

//float64保留小数点后两位
func Float64SaveAfterTwo(f float64) float64 {
	return Float64SavePointAfter(f, 2)
}
