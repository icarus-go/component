package service

import (
	"pmo-test4.yz-intelligence.com/kit/component/gins"
)

// User 用户业务逻辑
var User user

type user struct{}

func (*user) Get(ctx *gins.Context) (body interface{}, err error) {

	return
}
