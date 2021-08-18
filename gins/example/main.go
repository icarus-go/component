package main

import (
	"pmo-test4.yz-intelligence.com/kit/component/gins"

	// 加载 API 控制器
	"pmo-test4.yz-intelligence.com/kit/component/gins/example/arg"
	_ "pmo-test4.yz-intelligence.com/kit/component/gins/example/model"
	// 加载 Web 控制器
	_ "pmo-test4.yz-intelligence.com/kit/component/gins/example/controller"
	_ "pmo-test4.yz-intelligence.com/kit/component/gins/example/controller/web"
	// 加载 Web HTML 模板
	_ "pmo-test4.yz-intelligence.com/kit/component/gins/example/view"
)

func main() {

	arg.Config = &gins.Config{
		Name:    arg.Name,
		Host:    arg.Host,
		IP:      arg.Ip,
		Port:    arg.Port,
		Timeout: arg.Timeout,
		Debug:   arg.Debug,
		Pprof:   arg.Pprof,
		Version: VERSION,
	}

	// 静态文件

	// 使用Action风格路由全局中间件
	// gins.Use(gins.Action("/api"))

	// 重定向默认页
	//gins.GET("/", func(ctx *gins.Context) {
	//	//logger.Infof("默认页跳转：%s", ctx.Request.URL.Path)
	//	ctx.Redirect(http.StatusFound, "/web/app")
	//
	//	//logger.Info("默认页跳转-结束")
	//})

	gins.Use(func(ctx *gins.Context) {
		ctx.Next()
	})

	// 多分支多文件的

	// flag ->
	// 讀取參數 -> 獲取配置信息
	// 需要借助配置中心（etcd）
	// yaml
	// -> zap

	gins.AddInit(
		func() {
			//db.Init() // 初始化DB (放在第一位)
			//
			//err := consumer.Init() // 初始化MQ消费者 (假如有用到MQ)
			//if err != nil {
			//	panic(err)
			//}
			//
			//err = producer.Init() // 初始化MQ生产者 (假如有用到MQ)
			//if err != nil {
			//	panic(err)
			//}
			//
			//err = task.Init() // 启动任务 (假如有用到Task)
			//if err != nil {
			//	panic(err)
			//}
		},
	)

	// ./ 是指当前项目路径
	gins.Engine().Static("static", "./gins/example/static/dist") // 该路径取决于main文件存放在何处

	gins.Run(arg.Config)

	//task.Manager.Stop() // 退出任务

	//mq.Exit() // 退出MQ
}
