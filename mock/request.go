package mock

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
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

func (c *Request) AddParam(key, value string) *Request {
	c.values.Add(key, value)
	return c
}

func (c *Request) AddJSON(key string, value interface{}) *Request {
	c.json.Add(key, value)
	return c
}

func (c *Request) SetValues(parameter url.Values) *Request {
	c.values = parameter
	return c
}

func (c *Request) SetJSON(jsonObject *JSON) *Request {
	c.json = jsonObject
	return c
}

func (c *Request) AddHeader(key, value string) *Request {
	c.header.Add(key, value)
	return c
}

func (c *Request) Get(url string) (*Result, error) {
	if strings.Contains(url, "?") {
		url += "&"
	} else {
		url += "?"
	}

	url += c.values.Encode()

	c.AddHeader("content-Type", "application/x-www-form-urlencoded")

	return c.do("GET", url, nil)
}

func (c *Request) POST(url string) (*Result, error) {
	var data []byte

	if c.values != nil {
		data = []byte(c.values.Encode())
		c.AddHeader("content-Type", "application/x-www-form-urlencoded")
	}

	if c.json != nil {
		data = c.json.Body()
		c.AddHeader("content-Type", "application/json;charset=utf-8")
	}

	return c.do("POST", url, bytes.NewReader(data))
}

func (c *Request) do(method, url string, data *bytes.Reader) (*Result, error) {

	var req *http.Request
	var err error

	if data != nil {
		req, err = http.NewRequest(method, url, data)
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	// Header map[string][]string

	req.Header = c.header

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
