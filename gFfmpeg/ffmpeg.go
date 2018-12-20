package gFfmpeg

import (
	"github.com/lucky-lee/gutil/gCommand"
	"github.com/lucky-lee/gutil/gFile"
	"github.com/lucky-lee/gutil/gPath"
	"os/exec"
)

//amr -> mp3
func AmrToMp3(url string, fileName string) string {
	filePath := gPath.LogByName("ffmpeg/mp3")

	filePathName := filePath + fileName

	gFile.DirAutoCreate(filePath) //创建文件夹

	cmd := exec.Command("ffmpeg", "-i", url, filePathName)
	gCommand.Exec(cmd)

	return filePathName
}

//wav -> amr
func WavToAmr(url string, fileName string) string {
	filePath := gPath.LogByName("ffmpeg/amr")
	filePathName := filePath + fileName

	gFile.DirAutoCreate(filePath) //创建文件夹

	cmd := exec.Command("ffmpeg", "-i", url, "-ar", "8000", "-ab", "12.2k", "-ac", "1", filePathName)
	gCommand.Exec(cmd)

	return filePathName
}

//视频第一帧
func VideoFirstFrame(url string, fileName string) string {
	filePath := gPath.LogByName("ffmpeg/videoFrame")
	filePathName := filePath + fileName

	gFile.DirAutoCreate(filePath) //创建文件夹

	cmd := exec.Command("ffmpeg", "-i", url, "-y", "-f", "image2", "-ss", "2", "-vframes", "1", filePathName)
	gCommand.Exec(cmd)

	return filePathName
}
