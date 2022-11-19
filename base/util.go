package base

import (
	"log"
	"regexp"
)

func SetLoginHeader() map[string]string {
	header := make(map[string]string)

	header["Origin"] = "https://auth.dgut.edu.cn"
	header["Cache-Control"] = "max-age=0"
	header["Referer"] = "auth.dgut.edu.cn/authserver/login?service=https%3A%2F%2Fauth.dgut.edu.cn%2Fauthserver%2Foauth2.0%2FcallbackAuthorize%3Fclient_id%3D1021534300621787136%26redirect_uri%3Dhttps%253A%252F%252Fyqfk-daka.dgut.edu.cn%252Fnew_login%252Fdgut%26response_type%3Dcode%26client_name%3DCasOAuthClient"
	header["Content-Type"] = "application/x-www-form-urlencoded"

	return header
}

func SetAuthorizeHeader(cookie string) map[string]string {
	header := make(map[string]string)

	header["Origin"] = "https://auth.dgut.edu.cn"
	header["Cache-Control"] = "max-age=0"
	header["Referer"] = "auth.dgut.edu.cn/authserver/login?service=https%3A%2F%2Fauth.dgut.edu.cn%2Fauthserver%2Foauth2.0%2FcallbackAuthorize%3Fclient_id%3D1021534300621787136%26redirect_uri%3Dhttps%253A%252F%252Fyqfk-daka.dgut.edu.cn%252Fnew_login%252Fdgut%26response_type%3Dcode%26client_name%3DCasOAuthClient"
	header["Content-Type"] = "application/x-www-form-urlencoded"
	header["Cookie"] = cookie
	return header
}
func SetAuthHeader() map[string]string {
	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	header["Origin"] = "https://yqfk-daka.dgut.edu.cn"
	header["Referer"] = "https://yqfk-daka.dgut.edu.cn/"
	header["Host"] = "yqfk-daka-api.dgut.edu.cn"
	return header
}
func SetRecordHeader(token string) map[string]string {
	header := make(map[string]string)
	header["Host"] = "yqfk-daka-api.dgut.edu.cn"
	header["Origin"] = "https://yqfk-daka.dgut.edu.cn"
	header["Referer"] = "https://yqfk-daka.dgut.edu.cn/"
	header["Authorization"] = token
	return header
}

func SubMatchString(str, sub string) []string {
	compile := regexp.MustCompile(sub)
	if compile == nil {
		log.Fatal("err : PHP SESSION 正则出错...")
		return []string{}
	}
	subMatch := compile.FindStringSubmatch(str)
	return subMatch
}
