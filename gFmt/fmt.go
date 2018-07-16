package gFmt

import (
	"fmt"
	"github.com/lucky-lee/gutil/gTime"
)

func Println(a ...interface{}) {
	fmt.Println(gTime.StrNow(), a)
}
