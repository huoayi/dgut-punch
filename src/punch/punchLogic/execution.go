package punchLogic

import (
	"io/ioutil"
	"log"
	"net/http"
	"yqdk/base"
)

func GetExecuation() string {
	Client := http.Client{}
	url := "https://auth.dgut.edu.cn/authserver/login?service=https%3A%2F%2Fauth.dgut.edu.cn%2Fauthserver%2Foauth2.0%2FcallbackAuthorize%3Fclient_id%3D1021534300621787136%26redirect_uri%3Dhttps%253A%252F%252Fyqfk-daka.dgut.edu.cn%252Fnew_login%252Fdgut%26response_type%3Dcode%26client_name%3DCasOAuthClient"
	method := "GET"

	req, err := http.NewRequest(method, url, nil)

	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Add("Cache-Control", "max-age=0")
	req.Header.Add("Referer", "https://yqfk-daka.dgut.edu.cn/")
	req.Header.Add("Sec-Fetch-Dest", "document")
	req.Header.Add("sec-ch-ua", "\"Chromium\";v=\"106\", \"Google Chrome\";v=\"106\", \"Not;A=Brand\";v=\"99\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("Sec-Fetch-Mode", "navigate")
	req.Header.Add("Sec-Fetch-Site", "same-site")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", "apifox/1.0.0 (https://www.apifox.cn)")
	req.Header.Add("Host", "auth.dgut.edu.cn")
	req.Header.Add("Connection", "keep-alive")

	if err != nil {
		log.Fatal(err)
		return ""
	}
	res, err := Client.Do(req)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	subString := base.SubMatchString(string(body), `name="execution" value="(.*?)"`)
	execution := subString[len(subString)-1]
	return execution
}
