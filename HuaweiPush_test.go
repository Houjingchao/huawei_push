package huawei_push

import (
	"testing"
	"fmt"
	"github.com/Houjingchao/huawei_push/consts"
	"github.com/Houjingchao/huawei_push/model"
)

func TestGetToken(t *testing.T) {
	res := GetToken(consts.CLIENT_ID, consts.CLIENT_SECRECT)
	fmt.Println(res)
}

func TestPushByToken(t *testing.T) {
	palyod := model.NewHuaweiMessage().SetContent("huawei-content").SetTitle("huawei-title").JSON()
	fmt.Println(palyod)
	res := PushByToken("accesstoken",
		`list`,
		palyod)
	fmt.Println(res)
}
