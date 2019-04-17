package gFile

import (
	"github.com/lucky-lee/gutil/gFmt"
	"github.com/lucky-lee/gutil/gPath"
	"github.com/lucky-lee/gutil/gStr"
	"io"
	"net/http"
	"os"
)

//file download
func Download(url string, localPath string) string {
	gFmt.Println("下载文件地址:", url)

	if localPath == "" { //no file path use tmp path
		localPath = gPath.LogTmp()
	}

	DirAutoCreate(localPath) //auto create dir

	resp, err := http.Get(url)

	if err != nil {
		gFmt.Println("download file err:", err)
		return ""
	}

	defer resp.Body.Close()

	fileContentType := resp.Header.Get("Content-Type")

	gFmt.Println("Content-Type", fileContentType)

	if TypeByContentType(fileContentType) == "" { //not file
		return ""
	}

	fileName := DownloadFullName(url, fileContentType)
	fileLocal := localPath + fileName //local file path

	gFmt.Println("localFilePath:", fileLocal)
	gFmt.Println("statusCode", resp.StatusCode)

	if IsExist(fileLocal) {
		gFmt.Println("already exist file:", fileLocal)
		return fileName
	}

	if resp.StatusCode == http.StatusOK { //http code equal 200 and download
		out, err1 := os.Create(fileLocal)

		defer out.Close()

		if err1 == nil {
			io.Copy(out, resp.Body)
		} else {
			gFmt.Println("err1", err1)
		}
	}

	return fileName
}

//downloadfile full name
func DownloadFullName(url string, contentType string) string {
	return DownloadName(url) + TypeByContentType(contentType)
}

//dowanloadfile name
func DownloadName(url string) string {
	return gStr.Md5(url)
}
