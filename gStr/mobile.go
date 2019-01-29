package gStr

import "regexp"

var defReg = `^1([3578][0-9]|14[57]|5[^4])\d{8}$`

//设置-手机验证-正则
func SetMobileReg(s string) {
	defReg = s
}

//是否为手机号
func IsMobile(mobile string) bool {
	rgx := regexp.MustCompile(defReg)

	return rgx.MatchString(mobile)
}
