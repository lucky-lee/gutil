package gPath

import (
	"path/filepath"
	"os"
	"log"
	"strings"
	"fmt"
	"lucky/gutil"
)

//当前目录
func Current() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}

func Base() string {
	return fmt.Sprintf("%s/%s/", Current(), gutil.GetLogName())
}

func ByName(name string) string {
	return Base() + name + "/"
}

func Tmp() string {
	return ByName("tmp")
}

