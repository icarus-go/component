package gins

import (
	"fmt"
	selfConstant "github.com/icarus-go/component/constant"
	"github.com/icarus-go/data/params"
	"github.com/icarus-go/data/result"
	"github.com/icarus-go/data/result/constant"
	"net/http"
)

type api struct {
	ctx         *Context
	result      result.ApiResult
	rawResult   []byte
	filename    string
	contentType string
}

// SetError 设置错误信息
func (a *api) SetError(err error) {
	a.result.Msg = err.Error()

	if e, ok := err.(*APIError); ok {
		a.result.Code = e.code
		a.result.Data = e.data
		return
	}

	a.result.Code = constant.RESPONSE_ERROR
}

func (a *api) SetContentType(contentType selfConstant.ContentType) {
	a.contentType = contentType.Value()
}

// SetMsg 设置信息，code默认 RESPONSE_ERROR
func (a *api) SetMsg(msg string, code ...constant.ResponseType) {
	a.result.Msg = msg
	if len(code) == 1 {
		a.result.Code = code[0]
		return
	}

	a.result.Code = constant.RESPONSE_ERROR
}

//SetResponseOK
//  Author: Kevin·CC
//  Description: 空消息成功响应
func (a *api) SetResponseOK() {
	a.result.Code = constant.RESPONSE_OK
}

//SetOKMsg
//  Author: Kevin·CC
//  Description: OK 并且设置响应消息
//  Param msg 消息内容
func (a *api) SetOKMsg(msg string) {
	a.result.Code = constant.RESPONSE_OK
	a.result.Msg = msg
}

// SetData 设置输出的model
func (a *api) SetData(data interface{}) {
	a.result.Code = constant.RESPONSE_OK
	a.result.Data = data
}

//SetPageResult 设置分页返回
func (a *api) SetPageResult(list interface{}, total int64, page, pageSize int) {
	a.result.Code = constant.RESPONSE_OK
	a.result.Data = result.NewPageResult(list, total, page, pageSize)
}

//SetPaging 根据分页对象设置分页返回
func (a *api) SetPaging(list interface{}, total int64, paging params.Paging) {
	a.result.Code = constant.RESPONSE_OK
	a.result.Data = result.NewPageResult(list, total, paging.Page, paging.PageSize)
}

// SetDataKV 设置KV，会覆盖掉 SetData
func (a *api) SetDataKV(key string, value interface{}) {
	a.result.Code = constant.RESPONSE_OK
	if a.result.DataKV == nil {
		a.result.DataKV = make(map[string]interface{})
	}

	a.result.DataKV[key] = value
}

//SetRaw
//  Author: Kevin·CC
//  Description: 设置字节返回
//  Param rawResult
//  Param contentType
func (a *api) SetRaw(rawResult []byte, contentType selfConstant.ContentType) {
	a.rawResult = rawResult
	a.SetContentType(contentType)
}

//SetFile
//  Author: Kevin·CC
//  Description: 设置文件返回
//  Param raw
//  Param filename
func (a *api) SetFile(filename string, raw []byte) {
	a.filename = filename
	a.rawResult = raw
}

// SetRawResult 设置原始内容输出，content-Type为application/json，优先响应
func (a *api) SetRawResult(rawResult []byte, contentType ...selfConstant.ContentType) {
	a.rawResult = rawResult

	if len(contentType) != 0 {
		a.contentType = contentType[0].Value()
	}
}

// Render 立即渲染API
func (a *api) Render() {
	if a.result.DataKV != nil {
		a.result.Data = a.result.DataKV
	}

	if a.result.Data == nil {
		a.result.Data = struct{}{}
	}

	if a.rawResult != nil && a.filename == "" {
		a.ctx.Context.Data(http.StatusOK, a.contentType, a.rawResult)
		return
	}

	if a.filename != "" && a.rawResult != nil {
		a.ctx.Writer.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, a.filename))
		a.ctx.Data(http.StatusOK, selfConstant.FileStream.Value(), a.rawResult)
		return
	}

	a.ctx.Context.JSON(http.StatusOK, a.result)
}
