package core

import (
	"fmt"
	"net/http"
)

// 处理response
func (c *Context) Resp(code int, msg string, data interface{}, count int) *Context {
	c.resp = &response{
		Code:  code,
		Msg:   msg,
		Data:  data,
		Count: count,
	}

	return c
}

// 处理分页
func (c *Context) PageNate(page int, size int) *Context {
	c.resp.Page = page
	c.resp.PageSize = size

	return c
}

// json输出
func (c *Context) Json() {
	c.AbortWithStatusJSON(http.StatusOK, c.resp)
}

// 成功无数据输出
func (c *Context) SuccessNoData() {
	c.Resp(0, "ok", nil, 0).Json()
}

// 成功单数据输出
func (c *Context) SuccessOne(data interface{}) {
	c.Resp(0, "ok", data, 1).Json()
}

// 成功列表输出
func (c *Context) SuccessList(data interface{}, count int) {
	c.Resp(0, "ok", data, count).Json()
}

// 成功分页输出
func (c *Context) SuccessPage(data interface{}, count int, page int, size int) {
	c.Resp(0, "ok", data, count).PageNate(page, size).Json()
}

// 逻辑错误输出
func (c *Context) LogicErr(code int, msg string) {
	c.Resp(code, msg, nil, 0).Json()
}

// 系统异常输出
func (c *Context) SystemErr(errMsg string) {
	msg := "系统异常"
	if "" != errMsg {
		msg = fmt.Sprintf("系统异常: %s", errMsg)
	}

	c.Resp(100, msg, nil, 0).Json()
}

// 参数错误输出
func (c *Context) ParamErr(errMsg string) {
	msg := "数据错误"
	if "" != errMsg {
		msg = fmt.Sprintf("数据错误: %s", errMsg)
	}

	c.Resp(96, msg, nil, 0).Json()
}

// Token校验失败输出
func (c *Context) TokenAuthErr() {
	c.Resp(-100, "非法Token", nil, 0).Json()
}

// 生成Token失败输出
func (c *Context) TokenGenErr() {
	c.Resp(92, "生成Token失败", nil, 0).Json()
}

// 刷新Token失败输出
func (c *Context) TokenRefreshErr() {
	c.Resp(91, "刷新Token失败", nil, 0).Json()
}

// Ip白名单校验失败输出
func (c *Context) IpErr(ip string) {
	c.Resp(-200, "非法访问", map[string]string{"ip":ip}, 0).Json()
}

// 验证码校验失败输出
func (c *Context) VerifyCodeErr() {
	c.Resp(90, "无效的验证码", nil, 0).Json()
}

// 后台接口没有权限输出
func (c *Context) NoAccess() {
	c.Resp(-300, "没有权限", nil, 0).Json()
}
