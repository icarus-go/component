package mock

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"pmo-test4.yz-intelligence.com/kit/component/constant"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

type Request struct {
	values      url.Values
	json        *JSON
	header      http.Header
	contentType constant.ContentType
}

func NewRequest() *Request {
	return &Request{
		header: make(http.Header),
	}
}

func (r *Request) AddQuery(key, value string) *Request {
	if r.values == nil {
		r.values = make(url.Values)
	}

	r.values.Add(key, value)
	r.contentType = constant.URLEncode
	return r
}

func (r *Request) AddJSON(key string, value interface{}) *Request {
	if r.json == nil {
		r.json = NewJSON()
	}

	r.json.Add(key, value)
	r.contentType = constant.JSON

	return r
}

func (r *Request) AddFormData(key string, value []string) *Request {
	r.values[key] = value
	r.contentType = constant.FormData
	return r
}

func (r *Request) SetValues(parameter url.Values, contentType constant.ContentType) *Request {
	if r.values != nil {
		zap.L().Warn("将替换原有參數,请注意！", zap.String("body", string(r.json.Body())))
	}
	r.values = parameter
	r.contentType = contentType
	return r
}

func (r *Request) SetJSON(jsonObject *JSON) *Request {
	if r.json != nil {
		zap.L().Warn("将替换原有JSON,请注意！", zap.String("body", string(r.json.Body())))
	}

	r.json = jsonObject
	r.contentType = constant.JSON
	return r
}

func (r *Request) AddHeader(key, value string) *Request {
	r.header.Add(key, value)
	return r
}

func (r *Request) Get(url string) (*Result, error) {
	if strings.Contains(url, "?") {
		url += "&"
	} else {
		url += "?"
	}

	url += r.values.Encode()

	r.AddHeader("content-Type", r.contentType.Value())

	return r.do("GET", url)
}

func (r *Request) POST(url string) (*Result, error) {
	return r.do("POST", url)
}

func (r *Request) PUT(url string) (*Result, error) {
	return r.do("PUT", url)
}

func (r *Request) DELETE(url string) (*Result, error) {
	return r.do("DELETE", url)
}

func (r *Request) do(method, url string) (*Result, error) {
	var data []byte
	if r.values != nil {
		data = []byte(r.values.Encode())
	}

	if r.json != nil {
		data = r.json.Body()
	}

	if r.contentType == constant.FormData {
		r.AddHeader("Content-Length", strconv.Itoa(len(data)))
	}

	r.AddHeader("content-Type", r.contentType.Value())

	var req *http.Request
	var err error

	if data != nil {
		req, err = http.NewRequest(method, url, bytes.NewReader(data))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	// Header map[string][]string

	req.Header = r.header

	w := httptest.NewRecorder()

	Engine.ServeHTTP(w, req)

	if err != nil {
		err = errors.New("请求错误：" + err.Error())
		return nil, err
	}

	if w.Result().StatusCode != 200 {
		return nil, fmt.Errorf("请求错误：%d", w.Result().StatusCode)
	}

	return NewResult(w.Body.Bytes()), nil
}

func (r *Request) Reset() *Request {
	r.values = nil
	r.json = nil
	r.header = make(http.Header)
	return r
}
