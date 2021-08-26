package main

import (
	"pmo-test4.yz-intelligence.com/kit/component/gins"
	"pmo-test4.yz-intelligence.com/kit/component/gins/example/config"
	"pmo-test4.yz-intelligence.com/kit/component/gins/example/db"

	// 加载 API 控制器
	"pmo-test4.yz-intelligence.com/kit/component/gins/example/arg"
	_ "pmo-test4.yz-intelligence.com/kit/component/gins/example/model"
	// 加载 Web 控制器
	_ "pmo-test4.yz-intelligence.com/kit/component/gins/example/controller"
	_ "pmo-test4.yz-intelligence.com/kit/component/gins/example/controller/web"
)

func main() {

	arg.Config = &gins.Config{
		Name:          arg.Name,
		Host:          arg.Host,
		IP:            arg.Ip,
		Port:          arg.Port,
		Timeout:       arg.Timeout,
		Debug:         arg.Debug,
		Pprof:         arg.Pprof,
		IsCorsDisable: arg.IsCorsDisable,
		Version:       VERSION,
	}

	// 静态文件

	gins.Use(func(ctx *gins.Context) {
		ctx.Next()
	})

	gins.AddInit(
		func() {
			if err := config.Init(); err != nil {
				panic(err)
			}

			if err := db.Init(); err != nil {
				panic(err)
			}
		},
	)

	// ./ 是指当前项目路径
	gins.Engine().Static("static", "./gins/example/static/dist") // 该路径取决于main文件存放在何处

	gins.Run(arg.Config)

	//task.Manager.Stop() // 退出任务

	//mq.Exit() // 退出MQ
}
