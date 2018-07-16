package gStr

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
	"bytes"
)

//字符串截取
func Sub(str string, begin int, length int) string {
	rs := []rune(str)
	lth := len(rs)

	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}

	// 返回子串
	return string(rs[begin:end])
}

//合并字符串
func Merge(args ...string) (string) {
	b := bytes.Buffer{}

	for _, val := range args {
		b.WriteString(val)
	}
	return b.String()
}

//md5
func Md5(str string) string {
	data := []byte(str)
	md5Ctx := md5.New()
	md5Ctx.Write(data)
	sumStr := md5Ctx.Sum(nil)

	return hex.EncodeToString(sumStr)
}

//随机字符串
func Rand(length int, typeKey int) string {
	var str string
	if typeKey == 1 {
		str = "0123456789"
	} else {
		str = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	bs := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < length; i++ {
		result = append(result, bs[r.Intn(len(bs))])
	}

	return string(result)
}

//随机数
func RandNum(length int) string {
	return Rand(length, 1)
}
