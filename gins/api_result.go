package gins

import (
	"net/http"
	"pmo-test4.yz-intelligence.com/kit/component/apiconstant"
	"pmo-test4.yz-intelligence.com/kit/component/gins/result"
)

type api struct {
	ctx       *Context
	result    result.ApiResult
	rawResult []byte
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

//SetPageResult 设置分页返回
func (a *api) SetPageResult(list interface{}, total int64, page, pageSize int) {
	a.result.Code = apiconstant.RESPONSE_OK
	a.result.Data = result.NewPageResult(list, total, page, pageSize)
}

// SetDataKV 设置KV，会覆盖掉 SetData
func (a *api) SetDataKV(key string, value interface{}) {
	a.result.Code = apiconstant.RESPONSE_OK
	if a.result.DataKV == nil {
		a.result.DataKV = make(map[string]interface{})
	}

	a.result.DataKV[key] = value
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

	if a.result.DataKV != nil {
		a.result.Data = a.result.DataKV
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
