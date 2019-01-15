package gPath

import (
	"fmt"
	"github.com/lucky-lee/gutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//current path
func Current() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}

func CurrentFile(s string) string {
	return Current() + "/" + s
}

func BaseLog() string {
	return fmt.Sprintf("%s/%s/", Current(), gutil.GetLogName())
}

func LogByName(name string) string {
	return BaseLog() + name + "/"
}

func LogTmp() string {
	return LogByName("tmp")
}

func Static() string {
	return fmt.Sprintf("%s/static", Current())
}
