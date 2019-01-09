package gLog

import (
	"encoding/json"
	"fmt"
	"github.com/lucky-lee/gutil"
	"github.com/lucky-lee/gutil/gFile"
	"github.com/lucky-lee/gutil/gPath"
	"github.com/lucky-lee/gutil/gTime"
	"os"
)

func D(format string, value interface{}) {
	fileDir := gPath.LogByName("debug")
	toFile("Debug", fileDir, format, value)
}

func I(desc string, value interface{}) {
	fileDir := gPath.LogByName("info")
	toFile("Info", fileDir, desc, value)
}

func W(format string, value interface{}) {
	fileDir := gPath.LogByName("warn")
	toFile("Warn", fileDir, format, value)
}

func E(desc string, value interface{}) {
	fileDir := gPath.LogByName("err")
	toFile("Error", fileDir, desc, value)
}

func A(format string, value interface{}) {
	fileDir := gPath.LogByName("assert")
	toFile("Assert", fileDir, format, value)
}

func Json(desc string, val interface{}) {
	b, _ := json.Marshal(val)
	fileDir := gPath.LogByName("json")

	toFile(desc, fileDir, "json", string(b))
}

func Request(desc string, isJson bool, val interface{}) {
	fileDir := gPath.LogByName("request")

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

	fileDir := gPath.LogByName("sql")
	toFile(desc, fileDir, mess, nil)
}

func Log(desc string, format string, val interface{}) {
	fmt.Println(formatContent(desc, format, val))
}

//get log file name
func fileName() string {
	return fileLogName("")
}

func fileLogName(name string) string {
	if name == "" {
		name = gTime.StrDateToday()
	}

	return fmt.Sprintf("%s.log", name)
}

//get format log content
func formatContent(desc string, format string, val interface{}) string {
	var str string
	if val != nil {
		str = fmt.Sprintf("[%s-%s] %s | %s \n %v", gutil.GetLogName(), desc, gTime.StrNow(), format, val)
	} else {
		str = fmt.Sprintf("[%s-%s] %s | %s", gutil.GetLogName(), desc, gTime.StrNow(), format)
	}
	return str
}

//write content to log file
func toFile(name string, fileDir string, desc string, value interface{}) {
	//create dir
	gFile.DirAutoCreate(fileDir)

	//delete time out file
	delTimeOutFile(fileDir)

	status := gFile.Write(fileDir+fileName(), formatContent(name, desc, value))

	if !status {
		Log(name, desc, value)
	}
}

//time out date
var timeOutDate string

//delete time out log file
func delTimeOutFile(fileDir string) {
	toDate := gTime.StrDateDay(-gutil.GetLogDay())

	if timeOutDate == toDate { //if equal do not need excuse
		return
	}

	fileName := fileLogName(toDate)
	pathName := fileDir + fileName

	if gFile.IsExist(pathName) { //if file exist and remove it
		err := os.Remove(pathName)

		if err == nil {
			timeOutDate = toDate //set time out date
		}
	}
}
