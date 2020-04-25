package wechat

import (
	"base/config"
	"encoding/json"
	"strings"
)

var conf = config.GetConf().WeChat

func (w *WeCat) CodeToSession(code string) (*CodeToSessionResp, error) {
	var api = conf.Api.CodeToSession
	api = strings.ReplaceAll(api, "{appid}", w.AppId)
	api = strings.ReplaceAll(api, "{secret}", w.Secret)
	api = strings.ReplaceAll(api, "{code}", code)

	buf, err := w.setRequest(api, "GET", "", nil).doRequest()
	if nil != err {
		return nil, err
	}

	resp := new(CodeToSessionResp)
	err = json.Unmarshal(buf, resp)

	return resp, err
}
