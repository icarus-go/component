package gins

import (
	"net/http"
	"pmo-test4.yz-intelligence.com/kit/apiconstant"

	"github.com/gin-gonic/gin"
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
	//ctx.API.result.Code = apiconstant.RESPONSE_UNKNOW
	ctx.API.result.Msg = ""
	ctx.API.result.Data = nil
	ctx.API.result.dataKV = nil
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

type api struct {
	ctx       *Context
	result    apiResult
	rawResult []byte
}

type apiResult struct {
	Code   apiconstant.ResponseType `json:"code"`
	Msg    string                   `json:"msg"`
	Data   interface{}              `json:"data"`
	dataKV map[string]interface{}
}

type web struct {
	ctx *Context
}

// SetError 设置错误信息
func (a *api) SetError(err error) {
	a.result.Msg = err.Error()

	if e, ok := err.(*APIError); ok {
		a.result.Code = e.code
		a.result.Data = e.data
		return
	}

	a.result.Code = apiconstant.RESPONSE_ERROR
}

// SetMsg 设置信息，code默认 RESPONSE_ERROR
func (a *api) SetMsg(msg string, code ...apiconstant.ResponseType) {
	a.result.Msg = msg
	if len(code) == 1 {
		a.result.Code = code[0]
		return
	}

	a.result.Code = apiconstant.RESPONSE_ERROR
}

// SetData 设置输出的model
func (a *api) SetData(data interface{}) {
	a.result.Code = apiconstant.RESPONSE_OK
	a.result.Data = data
}

// SetDataKV 设置KV，会覆盖掉 SetData
func (a *api) SetDataKV(key string, value interface{}) {
	a.result.Code = apiconstant.RESPONSE_OK
	if a.result.dataKV == nil {
		a.result.dataKV = make(map[string]interface{})
	}

	a.result.dataKV[key] = value
}

// SetRawResult 设置原始内容输出，Content-Type为application/json，优先响应
func (a *api) SetRawResult(rawResult []byte) {
	a.rawResult = rawResult
}

func (a *api) json() {
	if a.rawResult != nil {
		a.ctx.Context.Data(http.StatusOK, "application/json;charset=utf-8", a.rawResult)
		return
	}

	if a.result.dataKV != nil {
		a.result.Data = a.result.dataKV
	}

	if a.result.Data == nil {
		a.result.Data = struct{}{}
	}

	a.ctx.Context.JSON(http.StatusOK, a.result)
}

// Render 立即渲染API
func (a *api) Render() {
	a.json()
}

// Render 立即渲染Web
func (w *web) Render(name string, obj interface{}) {
	w.ctx.Context.HTML(http.StatusOK, name, obj)
}
