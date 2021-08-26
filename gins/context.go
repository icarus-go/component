package gins

import (
	"github.com/gin-gonic/gin"
	"pmo-test4.yz-intelligence.com/kit/component/apiconstant"
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

// reset 重置Context
func (ctx *Context) reset(ginCtx *gin.Context) {
	ctx.Context = ginCtx
	// ctx.routerCtx, ctx.routerCancel = context.WithCancel(Instance.rootCtx)
	ctx.doneChan = make(chan struct{})

	ctx.stack = ""
	ctx.isAPI = false

	ctx.API.ctx = ctx
	ctx.API.result.Code = apiconstant.RESPONSE_UNKNOW
	ctx.API.result.Msg = ""
	ctx.API.result.Data = nil
	ctx.API.result.DataKV = nil
	ctx.API.rawResult = nil

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
