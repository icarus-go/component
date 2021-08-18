package api

import (
	"pmo-test4.yz-intelligence.com/kit/component/apiconstant"
	"pmo-test4.yz-intelligence.com/kit/component/gins"
)

// var apiRouter = gins.RouterGroup("/api").
// 	Use(gins.Action("/api")).
// 	Use(func(ctx *gins.Context) {
// 		ctx.SetIsAPI()
// 	})

// func init() {
// 	// 在Group使用中间件时，必须 路由路径 匹配时才会触发
// 	// Action路由命中时，会提前转走，不会进入下面流程
// 	apiRouter.GET("/", func(ctx *gins.Context) {
// 		ctx.API.SetMsg("找不到Action路由", apiconstant.RESPONSE_ERROR)
// 	})
// 	apiRouter.POST("/", func(ctx *gins.Context) {
// 		ctx.API.SetMsg("找不到Action路由", apiconstant.RESPONSE_ERROR)
// 	})
// }

func setResponseOK(ctx *gins.Context, msg ...string) {
	m := ""
	if len(msg) == 1 {
		m = msg[0]
	}

	ctx.API.SetMsg(m, apiconstant.RESPONSE_OK)
}
