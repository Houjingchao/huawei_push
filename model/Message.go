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

type ExtObj struct {
	Name string
}
type Ext struct {
	Action  string `json:"action"`
	Func    string `json:"func"`
	Collect string `json:"collect"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Url     string `json:"url"`
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

func (m *Message) SetIntent(intent string) *Message {
	m.Hps.Msg.Action.Param.Intent = intent
	return m
}

/**************************设置Param*************************************/

//Collect: "",
//Title: "",
//Content: "",
//Url: "",
func (m *Message) SetExtAction(Action string) *Message {
	m.Hps.Ext.Action = Action
	return m
}
func (m *Message) SetExtFunc(Func string) *Message {
	m.Hps.Ext.Func = Func
	return m
}
func (m *Message) SetExtCollect(collect string) *Message {
	m.Hps.Ext.Collect = collect
	return m
}
func (m *Message) SetExtTitle(title string) *Message {
	m.Hps.Ext.Title = title
	return m
}
func (m *Message) SetExtContent(content string) *Message {
	m.Hps.Ext.Collect = content
	return m
}

func (m *Message) SetExtUrl(url string) *Message {
	m.Hps.Ext.Url = url
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
						Intent: "",
					},
				},
			},
			Ext: Ext{ //扩展信息，含BI消息统计，特定展示风格，消息折叠。
				Action:  "",
				Func:    "",
				Collect: "",
				Title:   "",
				Content: "",
				Url:     "",
			},
		},
	}
}
