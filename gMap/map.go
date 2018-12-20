package gMap

import (
	"sort"
	"net/url"
	"github.com/lucky-lee/gutil"
	"github.com/lucky-lee/gutil/gStr"
)

//string val值的map排序
func MapSortKey(useMap map[string]interface{}) []gutil.StrBean {
	var keys []string
	var retSlice []gutil.StrBean

	for key, _ := range useMap {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, val := range keys {
		var bean gutil.StrBean

		bean.Key = val
		bean.Val = gStr.FormatAll(useMap[val], false) //useMap[val]

		retSlice = append(retSlice, bean)
	}

	return retSlice
}

//合并字符串map
func MapStrMerge(map1 map[string]string, map2 map[string]string) (map[string]string) {
	if len(map1) == 0 || len(map2) == 0 {
		return map[string]string{}
	}

	retMap := make(map[string]string)

	for key, val := range map1 {
		retMap[key] = val
	}

	for key, val := range map2 {
		retMap[key] = val
	}

	return retMap
}

//字符串map 转queryString
func MapStrToQueryStr(useMap map[string]interface{}) string {
	if useMap == nil {
		return ""
	}
	var retStr string
	for key, val := range useMap {
		retStr += key + "=" + url.QueryEscape(gStr.FormatAll(val, false)) + "&"
	}

	retStr = retStr[0 : len(retStr)-1]
	return retStr
}
