package code

import (
	"fmt"

	"pmo-test4.yz-intelligence.com/kit/component/apiconstant"
)

var message map[apiconstant.ResponseType]string

// ToInt 获取响应码
func ToInt(c apiconstant.ResponseType) int {
	return int(c)
}

// Msg 获取信息
func Msg(c apiconstant.ResponseType) string {

	if v, ok := message[c]; ok {
		return v
	}

	return ""
}

// MsgWithCode 获取错误码+信息 [10001]操作失败，类型不正确
func MsgWithCode(c apiconstant.ResponseType) string {

	if v, ok := message[c]; ok {
		return fmt.Sprintf("[%d]%s", c, v)
	}

	return "[90000]未知错误"
}
