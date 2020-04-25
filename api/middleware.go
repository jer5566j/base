package api

import (
	"base/api/core"
	"base/token"
	"base/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// TOKEN校验
func tokenAuth(context *core.Context, secret string, tokenData token.Token) {
	tokenStr := context.Request.Header.Get("Turnip")

	access := tokenData.Verify(tokenStr, secret)
	if access == false {
		utils.Log("error", fmt.Sprintf("TOKEN校验失败: %s %s", secret, tokenStr))
		context.TokenAuthErr()
		return
	}

	context.Set("token", tokenStr)
	context.Set("payload", tokenData)
	context.Next()
}

// 处理跨域
func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")

		var headerKeys []string
		for key, _ := range c.Request.Header {
			headerKeys = append(headerKeys, key)
		}

		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}

		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
			c.Header("Access-Control-Allow-Headers", "RRYB-Auth, Content-Length, X-CSRF-Token, Token,session, X_Requested_With, Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type, Expires, Last-Modified, Pragma, FooBar")
			c.Header("Access-Control-Max-Age", "172800")
			c.Header("Access-Control-Allow-Credentials", "false")
			//c.Set("content-type", "application/json")
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}

		c.Next()
	}
}
