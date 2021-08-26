package gins

import (
	"net/http"
	"pmo-test4.yz-intelligence.com/kit/component/apiconstant"
)

type api struct {
	ctx       *Context
	result    ApiResult
	rawResult []byte
}

type ApiResult struct {
	Code   apiconstant.ResponseType `json:"code"`
	Msg    string                   `json:"msg"`
	Data   interface{}              `json:"data"`
	dataKV map[string]interface{}   `json:"-"`
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
