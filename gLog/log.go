package gLog

import (
	"fmt"
	"lucky/gutil/gPath"
	"lucky/gutil/gTime"
	"lucky/gutil/gFile"
	"lucky/gutil"
	"encoding/json"
)

func D(format string, value interface{}) {
	fileDir := gPath.ByName("debug")
	toFile("Debug", fileDir, format, value)
}

func I(desc string, value interface{}) {
	fileDir := gPath.ByName("info")
	toFile("Info", fileDir, desc, value)
}

func W(format string, value interface{}) {
	fileDir := gPath.ByName("warn")
	toFile("Warn", fileDir, format, value)
}

func E(desc string, value interface{}) {
	fileDir := gPath.ByName("err")
	toFile("Error", fileDir, desc, value)
}

func A(format string, value interface{}) {
	fileDir := gPath.ByName("assert")
	toFile("Assert", fileDir, format, value)
}

func Json(desc string, val interface{}) {
	b, _ := json.Marshal(val)
	fileDir := gPath.ByName("json")

	toFile(desc, fileDir, "json", string(b))
}

func Request(desc string, isJson bool, val interface{}) {
	fileDir := gPath.ByName("request")

	if isJson {
		b, _ := json.Marshal(val)
		toFile(desc, fileDir, "json", string(b))
	} else {
		toFile(desc, fileDir, "json", val)
	}
}

func Sql(desc string, mess string) {
	if mess == "" {
		return
	}

	fileDir := gPath.ByName("sql")
	toFile(desc, fileDir, mess, nil)
}

func Log(desc string, format string, val interface{}) {
	fmt.Println(formatContent(desc, format, val))
}

//获取日志名称
func fileName() string {
	return fmt.Sprintf("%s.log", gTime.TsFormatYmd(0))
}

//日志-获取格式化后的内容
func formatContent(desc string, format string, val interface{}) string {
	var str string
	if val != nil {
		str = fmt.Sprintf("[%s-%s] %s | %s \n %s", gutil.GetLogName(), desc, gTime.StrNow(), format, val)
	} else {
		str = fmt.Sprintf("[%s-%s] %s | %s", gutil.GetLogName(), desc, gTime.StrNow(), format)
	}
	return str
}

//日志到file
func toFile(name string, fileDir string, desc string, value interface{}) {
	//创建文件夹
	gFile.DirAutoCreate(fileDir)
	status := gFile.Write(fileDir+fileName(), formatContent(name, desc, value))

	if !status {
		Log(name, desc, value)
	}
}
