package punch

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/textproto"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/golang/glog"
)

type Account struct {
	username string
	password string
}

func submit(username, psd string) {
	url := "https://yqfk-daka-api.dgut.edu.cn/record/"
	client := NewHttpClient()
	//   strings.NewReader(string(SetSubmitInfo(username,psd)))
	res, err := client.Post(url, SetRecordHeader(username, psd), strings.NewReader(string(SetSubmitInfo(username, psd))))
	if err != nil {
		glog.Errorf("submit error : %s", err)
		return
	}

	fmt.Println(res.HttpCode())
	fmt.Println(string(res.Data))
	data := make(map[string]string)
	_ = json.Unmarshal(res.Data, &data)
	fmt.Println(data["message"])

}

func SetSubmitInfo(username, psd string) []byte {
	info := GetRecordInfo(username, psd)
	infoData := make(map[string]interface{})
	json.Unmarshal(info, &infoData)
	subInfo := make(map[string]interface{})
	if infoData == nil {
		return nil
	}
	userData, i := infoData["user_data"].(map[string]interface{})
	if i == false {
		return nil
	}
	subInfo["submit_time"] = time.Now().Format("2006-01-02")
	subInfo["name"] = userData["name"]
	subInfo["faculty_name"] = userData["faculty_name"]
	subInfo["class_name"] = userData["class_name"]
	subInfo["username"] = userData["username"]
	subInfo["card_number"] = userData["card_number"]
	subInfo["identity_type"] = setDataWithInt(userData["identity_type"])
	subInfo["remark"] = userData["remark"]
	subInfo["tel"] = userData["tel"]
	subInfo["connect_person"] = userData["connect_person"]
	subInfo["connect_tel"] = userData["connect_tel"]
	subInfo["family_address_detail"] = userData["family_address_detail"]
	subInfo["current_country"] = userData["current_country"]
	subInfo["current_province"] = userData["current_province"]
	subInfo["current_city"] = userData["current_city"]
	subInfo["current_district"] = userData["current_district"]
	subInfo["latest_acid_test"] = userData["latest_acid_test"]
	subInfo["now_in_area_level"] = setDataWithInt(userData["now_in_area_level"])
	subInfo["current_in_city"] = setDataWithInt(userData["current_in_city"])
	subInfo["huji_region"] = setDataWithArray(userData["huji_region"])
	subInfo["family_region"] = setDataWithArray(userData["family_region"])
	subInfo["jiguan_region"] = userData["jiguan_region"]
	subInfo["current_region"] = setDataWithArray(userData["current_region"])
	subInfo["huji_region_name"] = userData["huji_region_name"]
	subInfo["family_region_name"] = userData["family_region_name"]
	subInfo["jiguan_region_name"] = userData["jiguan_region_name"]
	subInfo["card_type"] = userData["card_type"]
	subInfo["campus"] = setDataWithInt(userData["campus"])
	subInfo["have_diagnosis"] = setDataWithInt(userData["have_diagnosis"])
	subInfo["have_gone_important_area"] = setDataWithInt(subInfo["have_gone_important_area"])
	subInfo["have_contact_hubei_people"] = setDataWithInt(subInfo["have_contact_hubei_people"])
	subInfo["have_contact_illness_people"] = setDataWithInt(subInfo["have_contact_illness_people"])
	subInfo["have_isolation_in_dg"] = setDataWithInt(subInfo["have_isolation_in_dg"])
	subInfo["is_in_dg"] = setDataWithInt(subInfo["is_in_dg"])
	subInfo["is_new_in_dg"] = setDataWithInt(subInfo["is_new_in_dg"])
	subInfo["have_go_out"] = setDataWithInt(subInfo["have_go_out"])
	subInfo["is_specific_people"] = setDataWithInt(subInfo["is_specific_people"])
	subInfo["health_code_status"] = setDataWithInt(subInfo["health_code_status"])
	subInfo["in_controllerd_area"] = setDataWithInt(subInfo["in_controllerd_area"])
	subInfo["completed_vaccination"] = setDataWithInt(subInfo["completed_vaccination"])
	subInfo["have_stay_area"] = setDataWithInt(subInfo["have_stay_area"])
	subInfo["family_situation"] = setDataWithArray(subInfo["family_situation"])
	subInfo["health_situation"] = 1
	subInfo["body_temperature"] = "36.5"
	subInfo["is_in_school"] = 1
	subInfo["current_community_and_house"] = userData["current_community_and_house"]
	subInfo["current_street"] = userData["current_street"]
	subInfo["gps_address_name"] = userData["gps_address_name"]
	subInfo["gps_city"] = userData["gps_city"]
	subInfo["gps_city_name"] = userData["gps_city_name"]
	subInfo["gps_country"] = userData["gps_country"]
	subInfo["gps_country_name"] = userData["gps_country_name"]
	subInfo["gps_district"] = userData["gps_district"]
	subInfo["gps_district_name"] = userData["gps_district_name"]
	subInfo["gps_province"] = userData["gps_province"]
	subInfo["gps_province_name"] = userData["gps_province_name"]
	for k, v := range subInfo {
		fmt.Println(k, " ", v)
	}

	ans := make(map[string]interface{})
	ans["data"] = subInfo
	subByte, err := json.Marshal(ans)
	if err != nil {
		glog.Error(err)
		return nil
	}
	if subByte == nil {
		return nil
	}
	return subByte
}

func setDataWithInt(data interface{}) interface{} {
	if data == nil {
		return 0
	}
	return data
}
func setDataWithArray(data interface{}) interface{} {
	if data == nil {
		return []int{0}
	}
	return data
}

func GetRecordInfo(username, psd string) []byte {
	rul := "https://yqfk-daka-api.dgut.edu.cn/record/"
	client := NewHttpClient()

	map1 := SetRecordHeader(username, psd)
	// fmt.Println(map1["Authorization"])
	res, err := client.Get(rul, map1)
	if err != nil {
		glog.Errorf(err.Error())
		return nil
	}
	// fmt.Println(string(res.Data))
	return res.Data
}
func SetRecordHeader(username, psd string) map[string]string {
	header := make(map[string]string)
	header["Host"] = "yqfk-daka-api.dgut.edu.cn"
	header["Origin"] = "https://yqfk-daka.dgut.edu.cn"
	header["Referer"] = "https://yqfk-daka.dgut.edu.cn/"
	header["Authorization"] = "Bearer " + SetToken(username, psd)
	return header
}

func SetToken(username, psd string) string {
	s1 := GetCode(username, psd)
	s1 = `{"token":"` + s1 + `","state":"yqfk"}`

	url := "https://yqfk-daka-api.dgut.edu.cn/auth"
	client := NewHttpClient()

	resp, err := client.Post(url, nil, strings.NewReader(s1))
	if err != nil {
		glog.Errorf(err.Error())
		return ""
	}

	// fmt.Println(s2)

	k := make(map[string]string)
	json.Unmarshal(resp.Data, &k)

	return k["access_token"]
}

func GetCode(username, psd string) string {
	client := NewHttpClient()
	client.CheckRedirect(func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	})
	testUrl := "https://auth.dgut.edu.cn/authserver/oauth2.0/authorize?response_type=code&client_id=1021534300621787136&redirect_uri=https://yqfk-daka.dgut.edu.cn/new_login/dgut&state=yqfk"
	header := SetHeader()
	_, header["Cookie"] = SetCookieTest(username, psd)
	resp, err := client.Get(testUrl, header)
	if err != nil {
		glog.Error(err)
		return ""
	}
	loaction := resp.Headers.Get("Location")
	code := SubMatchString(loaction, `code=(.*?)&`)
	if code == nil {
		return ""
	}
	ans := code[0]
	return ans[5 : len(ans)-1]
}
func SetCookieTest(username, psd string) ([]*http.Cookie, string) {
	client := NewHttpClient()
	cookie := []*http.Cookie{}
	ckk := make(map[string]string)
	// golang http自动302跳转，但是不会在302时设置cookie
	// 所以如下代码取消自动跳转，然后手动设置cookie
	client.CheckRedirect(func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	})

	testUrl := "https://auth.dgut.edu.cn/authserver/login?service=https%3A%2F%2Fauth.dgut.edu.cn%2Fauthserver%2Foauth2.0%2FcallbackAuthorize%3Fclient_id%3D1021534300621787136%26redirect_uri%3Dhttps%253A%252F%252Fyqfk-daka.dgut.edu.cn%252Fnew_login%252Fdgut%26response_type%3Dcode%26client_name%3DCasOAuthClient"
	s := SetForm(Account{username, psd})
	s1 := GetPayload(s)
	payload := strings.NewReader(s1)
	resp, err := client.Post(testUrl, SetHeader(), payload)

	if err != nil {
		glog.Errorf("request error '%s', cause by %v", testUrl, err)
		return nil, ""
	}
	fmt.Println(username, psd)
Redirect:
	if resp.HttpCode() != 200 && resp.HttpCode() != 302 {
		glog.Errorf("request error[%d] %s ", resp.HttpCode(), testUrl)
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
		u, err := url.Parse(testUrl)
		if err != nil {
			glog.Errorf("parse host error '%s', cause by %v", u.Host, err)
			return nil, ""
		}
		cookie = setCookieWhenRedirect(client, u, resp, ckk)
		resp, _ = client.Get(location, nil)
		//可能多次跳转
		goto Redirect
	}
	var str string
	for k, v := range ckk {
		str = str + k + "=" + v + ";"
	}
	return cookie, str[:len(str)-1]
}

func setCookieWhenRedirect(client *HttpClient, domain *url.URL, resp *Response, cook map[string]string) []*http.Cookie {

	//get cookie (JSESSIONID=MDJkZDFkMjItZTc3OS00ZWY5LTkzZWUtODI3MDBmYWFkMzg2; path=/oauth; HttpOnly)
	cookieInHeader := resp.Headers[textproto.CanonicalMIMEHeaderKey("Set-Cookie")]

	var cookieStr string
	if len(cookieInHeader) != 0 {
		for _, i := range cookieInHeader {
			cookieStr = cookieStr + i + ";"
		}
	}
	if cookieStr == "" {
		return nil
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
	return cookies
}

type Response struct {
	Headers http.Header
	Data    []byte
}

func (resp *Response) HttpCodeIs200() (bool, error) {
	code := resp.HttpCode()
	if code != 200 {
		return false, fmt.Errorf("http-code: %d", code)
	}
	return true, nil
}

func (resp *Response) HttpCode() int {
	code := resp.Headers.Get("StatusCode")
	if code == "" {
		return 0
	}

	intCode, _ := strconv.Atoi(code)

	return intCode
}

type HttpClient struct {
	client *http.Client
}

func NewHttpClient() *HttpClient {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: tr,
	}

	return &HttpClient{client: client}
}

func (httpClient *HttpClient) CheckRedirect(fn func(req *http.Request, via []*http.Request) error) {
	httpClient.client.CheckRedirect = fn
}

func (httpClient *HttpClient) SetCookieJar(jar http.CookieJar) {
	httpClient.client.Jar = jar
}

func (httpClient *HttpClient) CookieJar() http.CookieJar {
	return httpClient.client.Jar
}

func (httpClient *HttpClient) Get(url string, header map[string]string) (*Response, error) {
	return httpClient.Call(http.MethodGet, url, header, nil)
}

func (httpClient *HttpClient) Post(url string, header map[string]string, body io.Reader) (*Response, error) {
	return httpClient.Call(http.MethodPost, url, header, body)
}

func (httpClient *HttpClient) Put(url string, header map[string]string, body io.Reader) (*Response, error) {
	return httpClient.Call(http.MethodPut, url, header, body)
}

func (httpClient *HttpClient) Delete(url string, header map[string]string) (*Response, error) {
	return httpClient.Call(http.MethodDelete, url, header, nil)
}

func (httpClient *HttpClient) Call(method, url string, header map[string]string, body io.Reader) (*Response, error) {
	req, err := http.NewRequest(method, url, body)

	if err != nil {
		return nil, err
	}

	if header != nil {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}

	resp, err := httpClient.client.Do(req)
	if err != nil {
		return nil, err
	}

	respHeader := resp.Header

	defer resp.Body.Close()

	respCode := resp.StatusCode
	respHeader.Set("StatusCode", strconv.Itoa(respCode))
	respHeader.Set("Status", resp.Status)

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read from response'body error. response code %v. response content: %v",
			respCode, string(content))
	}

	return &Response{Headers: respHeader, Data: content}, nil
}

func SetHeader() map[string]string {
	header := make(map[string]string)

	header["Origin"] = "https://auth.dgut.edu.cn"
	header["Cache-Control"] = "max-age=0"
	header["Referer"] = "auth.dgut.edu.cn/authserver/login?service=https%3A%2F%2Fauth.dgut.edu.cn%2Fauthserver%2Foauth2.0%2FcallbackAuthorize%3Fclient_id%3D1021534300621787136%26redirect_uri%3Dhttps%253A%252F%252Fyqfk-daka.dgut.edu.cn%252Fnew_login%252Fdgut%26response_type%3Dcode%26client_name%3DCasOAuthClient"
	header["Content-Type"] = "application/x-www-form-urlencoded"

	return header
}

func GetPayload(form map[string]string) string {
	var str string
	for k, v := range form {
		str = str + k + "=" + v + "&"
	}
	str = str[:len(str)-1]
	return str
}

func SetForm(account Account) map[string]string {
	data := make(map[string]string)
	data["username"] = account.username
	data["password"] = account.password
	data["captcha"] = ""
	data["_eventId"] = "submit"
	data["cllt"] = "userNameLogin"
	data["dllt"] = "generalLogin"
	data["lt"] = ""
	data["execution"] = GetExecuation()
	return data
}
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
	subString := SubMatchString(string(body), `name="execution" value="(.*?)"`)
	execution := subString[len(subString)-1]
	return execution
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
