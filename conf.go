package gutil

var logName = "gutil" //log file name

func SetLogName(value string) {
	logName = value
}

func GetLogName() string {
	return logName
}

var logDay = 5 //log file will keep 5 days

func SetLogDay(day int) {
	logDay = day
}

func GetLogDay() int {
	return logDay
}
