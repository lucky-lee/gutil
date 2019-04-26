package gTime

import (
	"fmt"
	"strconv"
	"time"
)

//time string covert timestamp
func StrToTs(timeStr string) (int64) {
	theTime, _ := time.ParseInLocation(STR_TIME_ALL, timeStr, Loc())
	return theTime.Unix()
}

//accurate timestamp string
func StrAccurate(ts int64) string {
	var retStr string
	nowTs := TsNow()
	offset := nowTs - ts

	if offset >= 0 { //以前
		if offset < 60 {
			retStr = strconv.FormatInt(offset, 10) + "秒前"
		} else if offset < 3600 {
			retStr = strconv.FormatInt(offset/60, 10) + "分钟前"
		} else if offset < 86400 {
			retStr = strconv.FormatInt(offset/3600, 10) + "小时前"
		} else if offset < 259200 {
			retStr = strconv.FormatInt(offset/86400, 10) + "天前"
		} else {
			retStr = TsFormatYmd(ts)
		}
	} else { //未来
		offset = ts - nowTs

		if offset < 60 {
			retStr = strconv.FormatInt(offset, 10) + "秒后"
		} else if offset < 3600 {
			retStr = strconv.FormatInt(offset/60, 10) + "分钟后"
		} else if offset < 86400 {
			retStr = strconv.FormatInt(offset/3600, 10) + "小时后"
		} else if offset < 259200 {
			retStr = strconv.FormatInt(offset/86400, 10) + "天后"
		} else {
			retStr = TsFormatYmd(ts)
		}

	}

	return retStr
}

//now time string
func StrNow() string {
	return time.Now().Format(STR_TIME_ALL)
}

//date string today
func StrDateToday() (dateStr string) {
	return StrDate(time.Now().Unix())
}

//date string today begin time
func StrDateTodayBegin() string {
	return fmt.Sprintf("%s 00:00:00", StrDateToday())
}

//date string today end time
func StrDateTodayEnd() string {
	return fmt.Sprintf("%s 23:59:59", StrDateToday())
}

//date string yesterday
func StrDateYesterday() (dateStr string) {
	return StrDateDay(-1)
}

//date string tomorrow
func StrDateTomorrow() (dateStr string) {
	return StrDateDay(1)
}

//date string use days
func StrDateDay(days int) (dateStr string) {
	ts := time.Now().AddDate(0, 0, days).Unix()
	return StrDate(ts)
}

//date string use timestamp
func StrDate(ts int64) (dateStr string) {
	return time.Unix(ts, 0).Format(STR_TIME_YMD)
}
