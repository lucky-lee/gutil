package gFfmpeg

import (
	"lucky/gutil/gPath"
	"os/exec"
	"lucky/gutil/gFile"
	"lucky/gutil/gCommand"
)

//amr -> mp3
func AmrToMp3(url string, name string) string {
	filePath := gPath.ByName("ffmpeg/mp3")

	filePathName := filePath + name

	gFile.DirAutoCreate(filePath) //创建文件夹

	cmd := exec.Command("ffmpeg", "-i", url, filePathName)
	gCommand.Exec(cmd)

	return filePathName
}

//wav -> amr
func WavToAmr(url string, name string) string {
	filePath := gPath.ByName("ffmpeg/amr")
	filePathName := filePath + name

	gFile.DirAutoCreate(filePath) //创建文件夹

	cmd := exec.Command("ffmpeg", "-i", url, "-ar", "8000", "-ab", "12.2k", "-ac", "1", filePathName)
	gCommand.Exec(cmd)

	return filePathName
}

//视频第一帧
func VideoFirstFrame(url string, name string) string {
	filePath := gPath.ByName("ffmpeg/videoFrame")
	filePathName := filePath + name

	gFile.DirAutoCreate(filePath) //创建文件夹

	cmd := exec.Command("ffmpeg", "-i", url, "-y", "-f", "image2", "-ss", "2", "-vframes", "1", filePathName)
	gCommand.Exec(cmd)

	return filePathName
}
