package notifyFromWechat

// Head the common fields of wechat notify
type Head struct {
	AppID          string `xml:"ToUserName"`
	FollowerOpenID string `xml:"FromUserName"`
	CreateTimeUnix int64  `xml:"CreateTime"`
	MsgType        string `xml:"MsgType"`
}
