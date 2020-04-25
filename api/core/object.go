package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 重写context
type Context struct {
	*gin.Context
	resp *response
}

// response结构体
type response struct {
	Code     int         `json:"rt_code"`
	Msg      string      `json:"msg"`
	Data     interface{} `json:"data"`
	Count    int         `json:"count"`
	Page     int         `json:"page,omitempty"`
	PageSize int         `json:"page_size,omitempty"`
}

// 自定义handler类型
type HandlerFunc func(c *Context)

// 自定义handler
func Handler(h HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := &Context{
			c,
			new(response),
		}

		h(ctx)
	}
}


