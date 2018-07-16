package gFile

import (
	"os"
	"fmt"
	"github.com/lucky-lee/gutil/gTime"
)

var fileMimeMap = map[string]string{
	"image/jpeg":               ".jpeg",
	"image/pjpeg":              ".jpeg",
	"image/jpg":                ".jpg",
	"image/png":                ".png",
	"image/gif":                ".gif",
	"image/webp":               ".webp",
	"image/bmp":                ".bmp",
	"video/mp4":                ".mp4",
	"application/octet-stream": ".jpg",
}

//文件类型-contentType获取
func TypeByContentType(typeStr string) string {
	fileType := fileMimeMap[typeStr]

	if fileType != "" {
		return fileType
	}
	return ""
}

//文件写入
func Write(filePath string, content string) bool {
	b := []byte(content + "\n")
	fd, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(gTime.StrNow(), "Error", "write.file.err", err)
		return false
	}

	fd.Write(b)
	fd.Close()

	return true
}

//file is exist
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	return false
}
