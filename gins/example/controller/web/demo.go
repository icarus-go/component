package web

import (
	"pmo-test4.yz-intelligence.com/kit/component/gins"
)

// Demo 测试模块
var Demo demo

type demo struct{}

func init() {
	webRouter.Group("/demo").
		GET("/panic", Demo.Panic)
}

// Panic 测试panic
func (*demo) Panic(ctx *gins.Context) {
	panic("测试panic")
}
