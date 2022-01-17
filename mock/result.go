package mock

import (
	"github.com/icarus-go/data/json"
	"github.com/icarus-go/data/result/constant"
)

type Result struct {
	Code constant.ResponseType `json:"code"`
	Msg  string                `json:"msg"`
	Data interface{}           `json:"data"`
	body []byte
}

func NewResult(body []byte) *Result {
	return &Result{body: body}
}

//Content 响应结果
func (r *Result) Content() string {
	return string(r.body)
}

//Unmarshal 序列化
func (r *Result) Unmarshal(target interface{}) error {
	return json.Unmarshal(r.body, &Result{
		Data: target,
	})
}
