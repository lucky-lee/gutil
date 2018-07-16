package gTime

import (
	"time"
	"strconv"
)

//时间字符串 转时间戳
func StrToTs(timeStr string) (int64) {
	theTime, _ := time.ParseInLocation(STR_TIME_ALL, timeStr, Loc())
	return theTime.Unix()
}

//时间字符串 精确
func StrAccurate(ts int64) string {
	var retStr string
	nowTs := TsNow()
	offset := nowTs - ts

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

	return retStr
}

//时间字符串 当前
func StrNow() string {
	return time.Unix(time.Now().Unix(), 0).Format(STR_TIME_ALL)
}
