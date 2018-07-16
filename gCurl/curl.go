package gCurl

import (
	"net/http"
	"strings"
	"io/ioutil"
	"github.com/lucky-lee/gutil/gMap"
	"github.com/lucky-lee/gutil/gLog"
)

//请求
func Request(method string, url string, useMap map[string]string) string {
	var reqUrl string

	reqUrl = url
	httpClient := &http.Client{}
	paramsQueryStr := gMap.MapStrToQueryStr(useMap)

	if method == "GET" {
		if paramsQueryStr != "" {
			reqUrl += "?" + paramsQueryStr
		}
	}

	paramsBody := strings.NewReader(paramsQueryStr)

	req, err := http.NewRequest(method, reqUrl, paramsBody)

	if err != nil {
		gLog.E("curl.err", err)
		return ""
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")

	resp, err := httpClient.Do(req)

	if err != nil {
		return ""
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return string(body)
	} else {
		return ""
	}
}

//响应状态码
func ResponseStatusCode(url string) int {
	if len(url) == 0 {
		return http.StatusNotFound
	}
	resp, err := http.Get(url)

	if err != nil {

		gLog.E("curlStatusCodeErr", err)
		return http.StatusNotFound
	}

	defer resp.Body.Close()

	return resp.StatusCode
}
