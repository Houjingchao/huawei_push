package huawei_push

import (
	"github.com/cocotyty/httpclient"
	"time"
	"strconv"
	"fmt"
	"github.com/Houjingchao/huawei_push/consts"
)

type HuaweiPush struct {
	ClientId     string
	ClientSecret string
}

//获取token  返回的expiretime 秒  当过期的时候
func (h HuaweiPush) GetToken() (string, error) {
	req, err := httpclient.
	Post(consts.ACCESSTOKENURL).
		Param("grant_type", consts.GRANTTYPE).
		Param("client_id", h.ClientId).
		Param("client_secret", h.ClientSecret).
		Send().String()
	if err != nil {
		return "", err
	}
	return req, nil
}

/**
access_token: 必选，使用OAuth2进行鉴权时的ACCESSTOKEN
nsp_ts: 必选，服务请求时间戳，自GMT 时间 1970-1-1 0:0:0至今的秒数。如果传入的时间与服务器时间相 差5分钟以上，服务器可能会拒绝请求。
nsp_svc: 必选， 本接口固定为openpush.message.api.send
device_token_list: 以半角逗号分隔的华为PUSHTOKEN列表，单次最多只是1000个
expire_time: 格式ISO 8601[6]:2013-06-03T17:30，采用本地时间精确到分钟
payload: 描述投递消息的JSON结构体，描述PUSH消息的:类型、内容、显示、点击动作、报表统计和扩展信 息。具体参考下面的详细说明。
 */
func(h HuaweiPush) PushByToken(access_token, device_token_list, payload string) error {
	now := time.Now()
	//dd, _ := time.ParseDuration("24h")
	//tomorrow := now.Add(dd)
	fmt.Println(strconv.Itoa(int(now.Unix())))
	_, err := httpclient.
	Post(consts.PUSHURL).
		Query("nsp_ctx", `{"ver":"1", "appId":"10781817"}`).
		Param("access_token", access_token). //access_token
		Param("nsp_svc", consts.NSP_SVC). //请求时间戳
		Param("nsp_ts", strconv.Itoa(int(now.Unix()))). //
		Param("device_token_list", device_token_list). //以半角逗号分隔的华为PUSHTOKEN列表，单次最多只是1000个
		Param("payload", payload).
		Param("expire_time", "2017-07-30T17:30").
		Send().String()

	if err != nil {
		return err
	}
	return err
}
