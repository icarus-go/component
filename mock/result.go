package mock

import (
	"encoding/json"

	"pmo-test4.yz-intelligence.com/kit/component/apiconstant"
)

type Result struct {
	Code apiconstant.ResponseType `json:"code"`
	Msg  string                   `json:"msg"`
	Data interface{}              `json:"data"`
	body []byte
}

func NewResult(body []byte) *Result {
	return &Result{body: body}
}

//Content 响应结果
func (r *Result) Content() string {
	return string(r.body)
}

func (r *Result) MarshalJSON(target interface{}) error {
	return json.Unmarshal(r.body, target)
}