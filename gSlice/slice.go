package gSlice

import (
	"github.com/lucky-lee/gutil"
	"net/url"
)

//create query string (strBean)
func QueryStrByStrBean(arr []gutil.StrBean) string {
	if len(arr) == 0 {
		return ""
	}

	var retStr string

	for _, val := range arr {
		retStr += val.Key + "=" + url.QueryEscape(val.Val) + "&"
	}

	retStr = retStr[0 : len(retStr)-1]
	return retStr
}

//slice reverse
func Reverse(s []interface{}) []interface{} {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}
