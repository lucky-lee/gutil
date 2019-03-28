package gCurl

import (
	"encoding/json"
	"encoding/xml"
	"github.com/lucky-lee/gutil/gLog"
	"github.com/lucky-lee/gutil/gMap"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	CONTENT_TYPE_FORM  uint8 = 0 //"application/x-www-form-urlencoded;charset=UTF-8"
	CONTENT_TYPE_JSON  uint8 = 1 //"application/json"
	CONTENT_TYPE_XML   uint8 = 2 //"application/xml"
	CONTENT_TYPE_EMPTY uint8 = 3 //""
)

var contentType uint8
var contentTypeStr string

//set content type json
func SetContentTypeJson() {
	contentType = CONTENT_TYPE_JSON
}

//set content type form
func SetContentTypeForm() {
	contentType = CONTENT_TYPE_FORM
}

//set content type form empty
func SetContentTypeEmpty() {
	contentType = CONTENT_TYPE_EMPTY
}

//set content type xml
func SetContentTypeXml() {
	contentType = CONTENT_TYPE_XML
}

func SetContentTypeStr() {
	if contentType == CONTENT_TYPE_JSON {
		contentTypeStr = "application/json;charset=UTF-8"
	} else if contentType == CONTENT_TYPE_FORM {
		contentTypeStr = "application/x-www-form-urlencoded;charset=UTF-8"
	} else if contentType == CONTENT_TYPE_XML {
		contentTypeStr = "application/xml;charset=UTF-8"
	} else if contentType == CONTENT_TYPE_EMPTY {
		contentTypeStr = ""
	}
}

//request
func Request(method string, url string, useMap map[string]interface{}) string {
	r := RequestReturnResponse(method, url, useMap)

	if r == nil {
		return ""
	}

	defer r.Body.Close()

	body, _ := ioutil.ReadAll(r.Body)
	return string(body)
}

//request xml
func RequestXml(method string, url string, params string) string {
	r := requestResponse(method, url, params)

	if r == nil {
		return ""
	}

	defer r.Body.Close()

	body, _ := ioutil.ReadAll(r.Body)

	return string(body)
}

//request and return response body
func RequestReturnResponse(method string, url string, useMap map[string]interface{}) *http.Response {
	var reqUrl string
	var paramsStr string

	reqUrl = url

	if contentType == CONTENT_TYPE_JSON {
		bs, _ := json.Marshal(useMap)
		paramsStr = string(bs)
	} else if contentType == CONTENT_TYPE_FORM {
		paramsStr = gMap.MapStrToQueryStr(useMap)
	} else if contentType == CONTENT_TYPE_EMPTY {
		paramsStr = gMap.MapStrToQueryStr(useMap)
	} else if contentType == CONTENT_TYPE_XML {
		x, _ := xml.Marshal(useMap)
		paramsStr = string(x)
	} else {
		return nil
	}

	if method == "GET" {
		if paramsStr != "" {
			reqUrl += "?" + paramsStr
		}
	}

	return requestResponse(method, url, paramsStr)
}

//请求-的返回数据
func requestResponse(method string, url string, params string) *http.Response {
	var reqUrl string

	reqUrl = url
	httpClient := &http.Client{}
	reader := strings.NewReader(params)
	req, err := http.NewRequest(method, reqUrl, reader)

	if err != nil {
		gLog.E("curl.err", err)
		return nil
	}

	//set content type string
	SetContentTypeStr()

	if contentTypeStr != "" {
		req.Header.Add("Content-Type", contentTypeStr)
	}

	//reset
	SetContentTypeForm()

	resp, err := httpClient.Do(req)

	if err != nil {
		gLog.E("requestRespErr", err)
		return nil
	}

	return resp
}

//get response status code
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
