package punchLogic

import (
	"encoding/json"
	"strings"
	"yqdk/base"

	"github.com/golang/glog"
)

func PostAuthGetBearerToken(code string) string {
	authResponseBody := buildAuthRequest(code)
	if authResponseBody == nil {
		return ""
	}
	// 从response里解析出token
	var authTokenStruct base.AuthAccessToken
	err := json.Unmarshal(authResponseBody, &authTokenStruct)
	if err != nil {
		glog.Errorf("auth token unmarshal error : %s", err.Error())
		return ""
	}
	bearerToken := "Bearer " + authTokenStruct.AccessToken

	return bearerToken
}
func buildAuthRequest(code string) []byte {
	// 构造请求去拿Bearer token
	authRequestBody := make(map[string]string)
	authRequestBody["token"] = code
	authRequestBody["state"] = "yqfk"
	jsonRequestBody, _ := json.Marshal(authRequestBody)
	client := base.NewHttpClient()
	authUrl := base.YQFK_DATA_API_URL + "auth"
	res, err := client.Post(authUrl, base.SetAuthHeader(), strings.NewReader(string(jsonRequestBody)))
	if err != nil {
		glog.Errorf("auth api error : %s", err.Error())
		return nil
	}
	// glog.Infof(string(res.Data))
	return res.Data
}
