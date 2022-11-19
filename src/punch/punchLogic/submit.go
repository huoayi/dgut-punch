package punchLogic

import (
	"encoding/json"
	"strings"
	"yqdk/base"

	"github.com/golang/glog"
)

func Submit(token string, info []byte) {
	url := base.YQFK_DATA_API_URL + "record/"
	client := base.NewHttpClient()
	name, subInfo := base.SetSubmitInfo(info)
	res, err := client.Post(url, base.SetRecordHeader(token), strings.NewReader(string(subInfo)))
	if err != nil {
		glog.Error(err)
	}
	data := make(map[string]string)
	_ = json.Unmarshal(res.Data, &data)
	glog.Infof("[%s] response msg : [%+v]", name, data["message"])
}
