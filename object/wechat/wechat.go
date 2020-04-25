package wechat

import (
	"base/config"
	"io/ioutil"
	"net/http"
	"strings"
)

type WeCat struct {
	AppId   string
	Secret  string
	Request Request
}

func InitWeChat() *WeCat {
	weChat := new(WeCat)
	weChat.AppId = config.GetConf().WeChat.AppId
	weChat.Secret = config.GetConf().WeChat.Secret

	return weChat
}

func (w *WeCat) setRequest(uri, method, contentType string, content interface{}, header ...RequestHeader) *WeCat {
	w.Request = *initRequest(uri, method).setHeader(header...).setContent(contentType, content)

	return w
}

func (w *WeCat) doRequest() ([]byte, error) {
	request, err := http.NewRequest(w.Request.Method, w.Request.Uri, strings.NewReader(w.Request.Content))
	if nil != err {
		return nil, err
	}

	if 0 < len(w.Request.Header) {
		for _, h := range w.Request.Header {
			for key, val := range h {
				request.Header.Add(key, val)
			}
		}
	}

	response, err := w.Request.Client.Do(request)
	if nil != err {
		return nil, err
	}

	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}
