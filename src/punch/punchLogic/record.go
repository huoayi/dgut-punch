package punchLogic

import (
	"github.com/golang/glog"
	"yqdk/base"
)

func PostGetRecordInfo(token string) []byte {
	url := base.YQFK_DATA_API_URL + "record/"
	client := base.NewHttpClient()
	res, err := client.Get(url, base.SetRecordHeader(token))
	if err != nil {
		glog.Error(err)
	}
	return res.Data
}
