package gTime

import "time"

//时间戳 当前
func TsNow() int64 {
	return time.Now().Unix()
}

//时间戳 转日期
func TsFormatYmd(unix int64) string {
	if unix == 0 {
		return time.Unix(time.Now().Unix(), 0).Format(STR_TIME_YMD)
	} else {
		return time.Unix(unix, 0).Format(STR_TIME_YMD)
	}
}

//时间戳 格式化 md
func TsFormatMd(ts int64) string {
	t, isNow := isNowYear(ts)

	if isNow {
		return t.Format(STR_TIME_MD)
	} else {
		return t.Format(STR_TIME_YMD)
	}
}

//时间戳 格式化 hi
func TsFormatHi(ts int64) string {
	t, isNow := isNowYear(ts)

	if isNow {
		return t.Format(STR_TIME_MD_HI)
	} else {
		return t.Format(STR_TIME_YMD_HI)
	}
}
