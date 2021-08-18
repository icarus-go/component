package errors

import (
	"pmo-test4.yz-intelligence.com/kit/component/apiconstant"
	"pmo-test4.yz-intelligence.com/kit/component/gins"
	"pmo-test4.yz-intelligence.com/kit/component/gins/example/errors/code"
)

/***** 90000,90500 用下面的方法 *****/

// New APIError API普通错误提示，90000
func New(msg string, code ...apiconstant.ResponseType) (ae *gins.APIError) {
	c := apiconstant.RESPONSE_ERROR
	if len(code) == 1 {
		c = code[0]
	}
	return gins.NewAPIError(msg, c)
}

// Except APIError API异常错误提示 90500
func Except(e error) (ae *gins.APIError) {
	return gins.NewAPIErrorWithLog("系统错误", e.Error())
}

/***** 自定义错误码用下面的方法 *****/

// LogicMask 自定义逻辑错误码 code：90000 msg: code对应的信息
func LogicMask(c apiconstant.ResponseType) (ae *gins.APIError) {
	return gins.NewAPIError(code.Msg(c), apiconstant.RESPONSE_ERROR)
}

// LogicMaskrWitdhCode 自定义逻辑错误码 code：90000 msg: [code]+code对应的信息
func LogicMaskrWitdhCode(c apiconstant.ResponseType) (ae *gins.APIError) {
	return gins.NewAPIError(code.MsgWithCode(c), apiconstant.RESPONSE_ERROR)
}

// Logic 自定义逻辑错误码 code：自定义错误码 msg: code对应的信息
func Logic(c apiconstant.ResponseType) (ae *gins.APIError) {
	return gins.NewAPIError(code.Msg(c), c)
}

// LogicWithCode 自定义逻辑错误码 code：自定义错误码 msg: [code]+code对应的信息
func LogicWithCode(c apiconstant.ResponseType) (ae *gins.APIError) {
	return gins.NewAPIError(code.MsgWithCode(c), c)
}
