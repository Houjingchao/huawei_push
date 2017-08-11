package model

import "encoding/json"

/**

{
"hps": {
"msg": {
"type": 3,
"body": {
"content": "Push message content",
"title": "Push message content"
},
"action": {
"type": 1,
"param": {
"intent": "#Intent;compo=com.rvr/.Activity;S.W=U;end"
}
}
},
"ext": {
"biTag": "Trump",
"icon": "http://upload.w.org/00/150pxsvg.png"
}
}
}
 */
type Message struct {
	Hps Hps `json:"hps"`
}

type Hps struct {
	Msg Msg `json:"msg"`
	Ext Ext `json:"ext"`
}
type Msg struct {
	Type   int `json:"type"`
	Body   Body `json:"body"`
	Action Action `json:"action"`
}
type Body struct {
	Content string `json:"content"`
	Title   string `json:"title"`
}
type Action struct {
	Type  int `json:"type"`
	Param Param `json:"param"`
}
type Param struct {
	Intent string `json:"intent"`
}

type Ext struct {
	BiTag string `json:"biTag"`
	Icon  string `json:"icon"`
}

func (m *Message) JSON() string {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func (m *Message) SetContent(content string) *Message {
	m.Hps.Msg.Body.Content = content
	return m
}

func (m *Message) SetTitle(title string) *Message {
	m.Hps.Msg.Body.Title = title
	return m
}

func NewHuaweiMessage() *Message {
	return &Message{
		Hps: Hps{
			Msg: Msg{
				Type: 3, //1  透传异步消息, 3 系统通知栏异步消息  注意:2和4以后为保留后续扩展使用
				Body: Body{
					Content: "",
					Title:   "",
				},
				Action: Action{
					Type: 1, //1 自定义行为，  2 打开URL ,3 打开App
					Param: Param{
						Intent: "#Intent;compo=com.rvr/.Activity;S.W=U;end",
					},
				},
			},
			Ext: Ext{ //扩展信息，含BI消息统计，特定展示风格，消息折叠。
				BiTag: "Trump",
				Icon:  "http://upload.w.org/00/150pxsvg.png",
			},
		},
	}
}
