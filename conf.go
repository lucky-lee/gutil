package gutil

var logName = "gutil" //日志名称

func SetLogName(value string) {
	logName = value
}

func GetLogName() string {
	return logName
}
