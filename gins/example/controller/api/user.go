package api

import (
	"pmo-test4.yz-intelligence.com/kit/component/apiconstant"
	"pmo-test4.yz-intelligence.com/kit/component/gins"
)

// User 用户 接口控制器
var User user

type user struct{}

// Get 获取指定用户信息
func (*user) Get(ctx *gins.Context) {
	ctx.API.SetMsg("", apiconstant.RESPONSE_OK)
}

// List 查询用户信息列表
func (*user) List(ctx *gins.Context) {
	ctx.API.SetMsg("暂不支持查询列表")
}
