package gTime

import (
	"time"
)

const (
	STR_TIME_ALL    = "2006-01-02 15:04:05"
	STR_TIME_YMD_HI = "2006-01-02 15:04"
	STR_TIME_MD_HI  = "01-02 15:04"
	STR_TIME_Y      = "2006"
	STR_TIME_YMD    = "2006-01-02"
	STR_TIME_MD     = "01-02"
	STR_TIME_HI     = "15:04"
	STR_TIME_HIS    = "15:04:05"
)

func Loc() *time.Location {
	loc, _ := time.LoadLocation("Local")
	return loc
}

//时间是否为今年
func isNowYear(ts int64) (retTime time.Time, isNow bool) {
	t := time.Unix(ts, 0)
	toNow := time.Unix(time.Now().Unix(), 0)
	nowYear := toNow.Format(STR_TIME_Y)
	timeYear := t.Format(STR_TIME_Y)

	if nowYear == timeYear {
		return t, true
	}
	return t, false
}
