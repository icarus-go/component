package gins

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/url"
	selfConstant "pmo-test4.yz-intelligence.com/kit/component/constant"

	"github.com/tidwall/gjson"
	"pmo-test4.yz-intelligence.com/kit/data/result/constant"
)

// Context 请求上下文
type Context struct {
	*gin.Context
	// routerCtx    context.Context
	// routerCancel context.CancelFunc
	doneChan chan struct{}

	stack string
	isAPI bool
	API   api
	Web   web
}

// Parse
//  Author: Kevin·CC
//  Description: 解析前端传递参数
//  Return jsonResult JSON结果
//  Return err 错误信息
func (ctx *Context) Parse() (jsonResult *gjson.Result, err error) {
	body, err := ctx.GetRawData()
	if err != nil {
		err = errors.New("json参数读取失败：" + err.Error())
		return
	}

	jsonResult = new(gjson.Result)
	*jsonResult = gjson.ParseBytes(body)

	if !jsonResult.IsObject() {
		err = errors.New("json参数格式错误")
	}

	return
}

// GetQuery
//  Author: Kevin·CC
//  Description: 获取前端传递的QUERY参数
//  Return url.Values URL 参数
func (ctx *Context) GetQuery() url.Values {
	return ctx.Request.URL.Query()
}

// reset 重置Context
func (ctx *Context) reset(ginCtx *gin.Context) {
	ctx.Context = ginCtx
	// ctx.routerCtx, ctx.routerCancel = context.WithCancel(Instance.rootCtx)
	ctx.doneChan = make(chan struct{})

	ctx.stack = ""
	ctx.isAPI = false

	ctx.API.ctx = ctx
	ctx.API.result.Code = constant.RESPONSE_UNKNOW
	ctx.API.result.Msg = ""
	ctx.API.result.Data = nil
	ctx.API.result.DataKV = nil
	ctx.API.rawResult = nil
	ctx.API.filename = ""
	ctx.API.contentType = selfConstant.JSON.Value()

	ctx.Web.ctx = ctx
}

// SetIsAPI 设置是否 API 请求标记
func (ctx *Context) SetIsAPI(args ...bool) {
	if len(args) == 0 {
		ctx.isAPI = true
		return
	}

	ctx.isAPI = args[0]
}

// IsAPI 是否 API 请求
func (ctx *Context) IsAPI() bool {
	return ctx.isAPI
}

// setPanic 设置异常信息
func (ctx *Context) setPanic(stack string) {
	ctx.stack = stack
}

// Panic 异常堆栈信息
// gins.On500 里可获取
func (ctx *Context) Panic() (stack string) {
	return ctx.stack
}
