package model

import "pmo-test4.yz-intelligence.com/kit/component/apiconstant"

type ApiResult struct {
	Code   apiconstant.ResponseType `json:"code"`
	Msg    string                   `json:"msg"`
	Data   interface{}              `json:"data"`
	DataKV map[string]interface{}   `json:"-"`
}
