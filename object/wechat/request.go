package wechat

import (
	"encoding/json"
	"net/http"
)

type Request struct {
	Uri     string
	Method  string
	Content string
	Header  []RequestHeader
	Client  *http.Client
}

type RequestHeader map[string]string

func initRequest(uri, method string) *Request {
	r := new(Request)
	r.Uri = uri
	r.Method = method
	r.Client = new(http.Client)

	return r
}

func (r *Request) setHeader(header ...RequestHeader) *Request {
	if nil != header {
		for _, h := range header {
			r.Header = append(r.Header, h)
		}
	}

	return r
}

func (r *Request) setContent(contentType string, content interface{}) *Request {
	switch contentType {
	case "json":
		buf, _ := json.Marshal(content)
		r.Header = append(r.Header, RequestHeader{"Content-Type": "application/json"})
		r.Content = string(buf)
		break
	}

	return r
}
