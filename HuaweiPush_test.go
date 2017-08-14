package huawei_push

import (
	"testing"
	"fmt"
	"github.com/Houjingchao/huawei_push/model"
)

func TestGetToken(t *testing.T) {
	huawei := HuaweiPush{
		ClientId:     "clientid",
		ClientSecret: "clientSecrtect",
	}
	res, _ := huawei.GetToken()
	fmt.Println(res)
}

func TestPushByToken(t *testing.T) {
	huawei := HuaweiPush{
		ClientId:     "clientid",
		ClientSecret: "clientSecrtect",
	}
	palyod := model.NewHuaweiMessage().SetContent("huawei-content").SetTitle("huawei-title").JSON()
	fmt.Println(palyod)
	res := huawei.PushByToken("accesstoken",
		`list`,
		palyod)
	fmt.Println(res)
}
