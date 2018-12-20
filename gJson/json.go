package gJson

import (
	"encoding/json"
	"github.com/lucky-lee/gutil/gFmt"
	"net/http"
)

//json encode
func Encode(i interface{}) string {
	b, err := json.Marshal(i)

	if err != nil {
		gFmt.Println("util.getJsonStr.error", err)
		return ""
	}

	return string(b)
}

//json decode
func Decode(s string, i interface{}) {
	err := json.Unmarshal([]byte(s), i)

	if err != nil {
		gFmt.Println("json.decode.err", err)
		gFmt.Println("json.decode.string", s)
	}
}

//write json
func Write(code int, w http.ResponseWriter, obj interface{}) {
	header := w.Header()

	w.WriteHeader(code)

	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = []string{"application/json; charset=utf-8"}
	}

	b, err := json.Marshal(obj)

	if err != nil {
		gFmt.Println("util.getJsonStr.error", err)
		return
	}

	w.Write(b)
}
