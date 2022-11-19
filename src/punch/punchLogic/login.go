package punchLogic

import (
	"net/http"
	"net/http/cookiejar"
	"net/textproto"
	"net/url"
	"strings"
	"yqdk/base"

	"github.com/golang/glog"
)

func GetCode(dataMap map[string]string) string {
	client := base.NewHttpClient()
	client.CheckRedirect(func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	})
	url := base.AUTHORIZE_URL
	cookie := GetCookie(dataMap)
	header := base.SetAuthorizeHeader(cookie)
	res, err := client.Get(url, header)
	if err != nil {
		glog.Errorf("GetCode Error msg: %s", err.Error())
		return ""
	}
	location := res.Headers.Get("Location")
	code := base.SubMatchString(location, `code=(.*?)&`)
	if len(code) == 0 {
		glog.Errorf("GetCode Error msg: code is not found")
		return ""
	}
	ans := code[0]
	return ans[5 : len(ans)-1]
}

func GetCookie(dataMap map[string]string) string {
	requestLoginHost, cookie := buildRequest(dataMap)
	if requestLoginHost == nil {
		return ""
	}
	ok, err := requestLoginHost.HttpCodeIs200()
	if err != nil {
		glog.Error("func requestLoginHost get problem:", err)
	} else if !ok {
		glog.Error("request fail:the code is", requestLoginHost.HttpCode())
	}
	return cookie
}

func buildRequest(dataMap map[string]string) (*base.Response, string) {
	client := base.NewHttpClient()
	loginUrl := base.LOGIN_URL
	//cookie := []*http.Cookie{}
	ckk := make(map[string]string)
	client.CheckRedirect(func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	})
	form := getPayload(dataMap)
	payload := strings.NewReader(form)
	resp, err := client.Post(loginUrl, base.SetLoginHeader(), payload)
	if err != nil {
		glog.Error(err.Error())
		return nil, ""
	}

Redirect:
	if resp.HttpCode() != 200 && resp.HttpCode() != 302 {
		glog.Errorf("request error[%d] %s ", resp.HttpCode(), loginUrl)
		return nil, ""
	}

	if resp.HttpCode() == 200 {
		//最后请求成功

	} else if resp.HttpCode() == 302 {
		// xWiki /oauth/login
		location := resp.Headers.Get("Location")
		if location == "" {
			location = resp.Headers.Get("location")
		}
		u, err := url.Parse(loginUrl)
		if err != nil {
			glog.Errorf("parse host error '%s', cause by %v", u.Host, err)
			return nil, ""
		}
		setCookieWhenRedirect(client, u, resp, ckk)
		resp, _ = client.Get(location, nil)
		//可能多次跳转
		goto Redirect
	}
	var str string
	for k, v := range ckk {
		str = str + k + "=" + v + ";"
	}
	return resp, str[:len(str)-1]
}

func setCookieWhenRedirect(client *base.HttpClient, domain *url.URL, resp *base.Response, cook map[string]string) {
	cookieInHeader := resp.Headers[textproto.CanonicalMIMEHeaderKey("Set-Cookie")]
	var cookieStr string
	if len(cookieInHeader) != 0 {
		for _, i := range cookieInHeader {
			cookieStr = cookieStr + i + ";"
		}
	}
	if cookieStr == "" {
		return
	}
	//build cookie

	cookies := []*http.Cookie{}

	parts := strings.Split(cookieStr, ";")
	for _, item := range parts {
		cookie := &http.Cookie{}
		if item == "" {
			continue
		}
		item = strings.Trim(item, " ")
		if item == "HttpOnly" {
			cookie.HttpOnly = true
			continue
		}
		subItems := strings.Split(item, "=")
		if len(subItems) != 2 {
			continue
		}
		if subItems[0] == "Path" || subItems[0] == "path" {
			cookie.Path = subItems[1]
			continue
		}
		cookie.Name = subItems[0]
		cookie.Value = subItems[1]
		cookies = append(cookies, cookie)
		cook[subItems[0]] = subItems[1]
	}
	// set cookie
	jar := client.CookieJar()
	if jar == nil {
		jar, _ = cookiejar.New(nil)
	}
	jar.SetCookies(domain, cookies)
	client.SetCookieJar(jar)
	// 设置成功后下次请求domain就会带上cookie

}

func getPayload(form map[string]string) string {
	var str string
	for k, v := range form {
		str = str + k + "=" + v + "&"
	}
	str = str[:len(str)-1]
	return str
}
