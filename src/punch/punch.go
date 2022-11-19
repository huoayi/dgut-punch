package punch

import (
	"github.com/golang/glog"
	"sync"
	"time"
	"yqdk/base"
	it "yqdk/src/global"
	pl "yqdk/src/punch/punchLogic"
)

func Start() {
	var FuncNum = it.Config.FuncMAX
	var userChannel = make(chan base.UserAccount, FuncNum)
	var dataChannel = make(chan base.DataForm, FuncNum)
	var codeChannel = make(chan string, FuncNum)
	var bearerTokenChannel = make(chan string, FuncNum)
	var threshold int
	var cnt int
	var wg sync.WaitGroup

	go func() {
		threshold += initData(userChannel)
	}()
	glog.Infof("start punch \n")
LOOP:
	for {
		select {
		case user, ok := <-userChannel:
			{
				if !ok {
					continue
				}
				go func(account base.UserAccount) {
					data := pl.BuildRequestData(account)
					var dc = base.DataForm{Data: data}
					dataChannel <- dc
				}(user)
			}
		case data := <-dataChannel:
			{
				go func(dc base.DataForm) {
					code := pl.GetCode(dc.Data)
					if code == "" {
						threshold--
						return
					}
					codeChannel <- code
				}(data)
			}
		case code := <-codeChannel:
			{
				go func(tk string) {
					bearerToken := pl.PostAuthGetBearerToken(tk)
					if bearerToken == "" {
						threshold--
						return
					}
					bearerTokenChannel <- bearerToken
				}(code)
			}
		case token := <-bearerTokenChannel:
			{
				wg.Add(1)
				go func(tk string) {
					// []byte
					recordInfo := pl.PostGetRecordInfo(tk)
					if recordInfo == nil {
						wg.Done()
						threshold--
						return
					}
					pl.Submit(tk, recordInfo)
					wg.Done()
				}(token)
				// when all user information is submitted, it will exit the infinite loop.
				cnt++
				if cnt == threshold {
					break LOOP
				}
			}
		}
	}
	wg.Wait()
	time.Sleep(time.Duration(1) * time.Second)
	glog.Infof("punch end, total:[%d]...", threshold)

}

func initData(ch chan base.UserAccount) int {
	ca := it.Config
	var u []base.UserAccount
	if ca.DataSource == 0 {
		u = ca.UserAccount
	} else if ca.DataSource == 1 {
		u = pl.LoadDataBaseConfig()
	}
	var userNum = len(u)
	if userNum == 0 {
		glog.Errorf("load data error...")
		return 0
	}
	glog.Infof("read file success")
	for _, row := range u {
		ch <- row
	}
	time.Sleep(time.Duration(1) * time.Second)
	close(ch)
	return userNum
}
