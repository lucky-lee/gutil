package gFile

import (
	"net/http"
	"os"
	"io"
	"github.com/lucky-lee/gutil/gFmt"
	"github.com/lucky-lee/gutil/gPath"
	"github.com/lucky-lee/gutil/gStr"
)

//文件下载
func Download(url string, localPath string) string {
	gFmt.Println("下载文件地址:", url)
	var filePath string //文件路径

	if localPath == "" { //没有指定path 在tmp文件下
		filePath = gPath.Tmp()
		DirAutoCreate(gPath.Tmp())
	} else {
		filePath = localPath
	}

	//下载图片保存到本地
	resp, err := http.Get(url)

	if err != nil {
		gFmt.Println("下载文件错误,err:", err)
		return ""
	}

	fileContentType := resp.Header.Get("Content-Type")

	gFmt.Println("Content-Type", fileContentType)

	if TypeByContentType(fileContentType) == "" { //不是文件
		defer resp.Body.Close()
		return ""
	}

	fileName := DownloadFullName(url, fileContentType)
	fileLocal := filePath + fileName //本地文件path

	gFmt.Println("localFilePath:", fileLocal)
	gFmt.Println("statusCode", resp.StatusCode)

	if IsExist(fileLocal) {
		gFmt.Println("已经存在file:", fileLocal)
		defer resp.Body.Close()
		return ""
	}

	if resp.StatusCode == http.StatusOK { //200下载
		out, err1 := os.Create(fileLocal)

		defer out.Close()

		if err1 == nil {
			io.Copy(out, resp.Body)
		} else {
			gFmt.Println("err1", err1)
		}
	}

	defer resp.Body.Close()

	return fileName
}

//下载文件-全名
func DownloadFullName(url string, contentType string) string {
	return DownloadName(url) + TypeByContentType(contentType)
}

//下载文件-名称
func DownloadName(url string) string {
	return gStr.Md5(url)
}
