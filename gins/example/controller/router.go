package controller

import (
	"pmo-test4.yz-intelligence.com/kit/component/dashboard"
	"pmo-test4.yz-intelligence.com/kit/component/gins/example/arg"
	"pmo-test4.yz-intelligence.com/kit/component/gins/example/controller/api"

	"pmo-test4.yz-intelligence.com/kit/component/gins"
)

var apiRouter = gins.Instance.Router.Group("/api").Use(func(context *gins.Context) {
	context.SetIsAPI()
})

func init() {

	apiRouter.GET("/info", func(context *gins.Context) {
		bill, _ := dashboard.New().Service.Info.ServiceName(arg.Name).Version(arg.Config.Version).Bill()
		context.API.SetData(bill)
	})

	apiRouter.Group("/demo").
		POST("/add", api.Demo.Add).
		GET("/get", api.Demo.Get).
		GET("/timeout", api.Demo.Timeout).
		GET("/panic", api.Demo.Panic).
		// Action风格路由
		POST("/update", api.Demo.Update).
		POST("/list", api.Demo.List)

	apiRouter.Group("/user").
		//GET("/", api.User.Get).
		GET("/list", api.User.List)

	//apiRouter.GET("/", func(ctx *gins.Context) {
	//	ctx.API.SetMsg("找不到Action路由", apiconstant.RESPONSE_ERROR)
	//})
	//apiRouter.POST("/", func(ctx *gins.Context) {
	//	ctx.API.SetMsg("找不到Action路由", apiconstant.RESPONSE_ERROR)
	//})

}
