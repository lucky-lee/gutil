package gSlice

import (
	"lucky/gutil"
	"net/url"
)

//生成query string (strBean)
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
