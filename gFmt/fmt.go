package gFmt

import (
	"fmt"
	"lucky/gutil/gTime"
)

func Println(a ...interface{}) {
	fmt.Println(gTime.StrNow(), a)
}
