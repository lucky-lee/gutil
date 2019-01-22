package gStr

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

//字符串截取
//str 需要截取的字符串
//begin 开始的位置
//length 需要截取的长度
func Sub(str string, begin int, length int) string {
	rs := []rune(str)
	lth := len(rs)

	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}

	// 超过最大长度
	if begin >= lth {
		begin = lth
	}

	//长度为0 截取到结尾
	if length == 0 {
		return string(rs[begin:])
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

//rand num
//use you want num create it
func RandNum(length int) string {
	return Rand(length, 1)
}

func RandStr(length int) string {
	return Rand(length, 0)
}

//format any one
func FormatAny(value interface{}) string {
	return formatAny(reflect.ValueOf(value), true)
}

func FormatAll(v interface{}, strQuote bool) string {
	return formatAny(reflect.ValueOf(v), strQuote)
}

//format type any one to string
func formatAny(value reflect.Value, strQuote bool) string {
	switch value.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(value.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(value.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(value.Float(), 'E', -1, 64)
	case reflect.Bool:
		return strconv.FormatBool(value.Bool())
	case reflect.String:
		if strQuote {
			return strconv.Quote(value.String())
		} else {
			return value.String()
		}
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return value.Type().String() + " 0x" +
			strconv.FormatUint(uint64(value.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return value.Type().String() + " value"
	}
}
