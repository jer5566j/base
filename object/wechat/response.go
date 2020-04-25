package wechat

type CodeToSessionResp struct {
	OpenId     string `json:"openid"`
	UnionId    string `json:"unionid"`
	SessionKey string `json:"session_key"`
	ErrCode    int64  `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}
