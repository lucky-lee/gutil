package gFile

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

//文件转base64
func FileBase64Encode(path string) (base64Str string) {
	buffFile, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println("read.file.error", err)
		return
	}

	//buff := make([]byte, 500000)
	base64Str = base64.StdEncoding.EncodeToString(buffFile)

	return
}

//base64 转文件
func FileBase64Decode(code, dest string) error {
	buff, err := base64.StdEncoding.DecodeString(code)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(dest, buff, 07440); err != nil {
		return err
	}

	return nil
}
