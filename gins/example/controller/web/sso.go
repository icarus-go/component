package web

import (
	"pmo-test4.yz-intelligence.com/kit/component/gins"
)

// SSO 登录控制器
var SSO sso

type sso struct{}

func init() {
	webRouter.Group("/sso").
		GET("/oauth", SSO.OAuth)
}

// OAuth 授权登录
func (*sso) OAuth(ctx *gins.Context) {
	ctx.Web.Render("sso_oauth", nil)
}
