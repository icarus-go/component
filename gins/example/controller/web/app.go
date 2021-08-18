package web

import (
	"pmo-test4.yz-intelligence.com/kit/component/gins"
)

// App SPA主页控制器
var App app

type app struct{}

func init() {
	webRouter.Group("/app").
		GET("/*action", App.SPA)
}

// SPA 单页应用
func (*app) SPA(ctx *gins.Context) {
	ctx.Web.Render("app", nil)
}
