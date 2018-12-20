package gCurl

import (
	"fmt"
	"github.com/lucky-lee/gutil/gStr"
	"strings"
)

//build http params query
func HttpBuildQuery(params map[string]interface{}) string {
	var arr []string

	for key, val := range params {
		str := fmt.Sprintf("%s=%s", key, gStr.FormatAll(val, false))
		arr = append(arr, str)
	}

	return strings.Join(arr, "&")
}
