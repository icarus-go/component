package mock

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"

	"go.uber.org/zap"
)

type Request struct {
	values url.Values
	json   *JSON
	header http.Header
}

func NewRequest() *Request {
	return &Request{
		header: make(http.Header),
	}
}

func (r *Request) AddParam(key, value string) *Request {
	r.values.Add(key, value)
	return r
}

func (r *Request) AddJSON(key string, value interface{}) *Request {
	if r.json == nil {
		r.json = NewJSON()
	}

	r.json.Add(key, value)
	return r
}

func (r *Request) SetValues(parameter url.Values) *Request {
	r.values = parameter
	return r
}

func (r *Request) SetJSON(jsonObject *JSON) *Request {
	if r.json != nil {
		zap.L().Warn("将替换原有JSON,请注意！", zap.String("body", string(r.json.Body())))
	}

	r.json = jsonObject

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

	r.AddHeader("content-Type", "application/x-www-form-urlencoded")

	return r.do("GET", url, nil)
}

func (r *Request) POST(url string) (*Result, error) {
	var data []byte

	if r.values != nil {
		data = []byte(r.values.Encode())
		r.AddHeader("content-Type", "application/x-www-form-urlencoded")
	}

	if r.json != nil {
		data = r.json.Body()
		r.AddHeader("content-Type", "application/json;charset=utf-8")
	}

	return r.do("POST", url, bytes.NewReader(data))
}

func (r *Request) do(method, url string, data *bytes.Reader) (*Result, error) {

	var req *http.Request
	var err error

	if data != nil {
		req, err = http.NewRequest(method, url, data)
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
