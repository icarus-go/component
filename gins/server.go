package gins

import (
	"context"
	"html/template"
	"net/http"
	"net/http/pprof"
	"time"

	"github.com/gin-gonic/gin"

)

// InitFunc 安全初始化函数
type InitFunc func()

// Server 服务器对象
type Server struct {
	engine       *gin.Engine        // gin Engine
	Middleware   *Middleware        // 全局中间件
	Router       *Router            // gin Router 封装
	templ        *template.Template // 模板资源
	initFuncList []InitFunc         // 安全初始化函数列表

	server     *http.Server // http服务器
	rootCtx    context.Context
	rootCancel context.CancelFunc

	on404 HandlerFunc
	on500 HandlerFunc

	Config *Config
}

// New 创建新的GinServer实例
func New() (gs *Server, err error) {
	gs = &Server{
		engine: gin.New(),
	}

	// 封闭自定义 Middleware ，全局
	gs.Middleware = &Middleware{engine: gs.engine}

	// 封装自定义 Router
	gs.Router = &Router{RouterGroup: &gs.engine.RouterGroup}

	gs.rootCtx, gs.rootCancel = context.WithCancel(context.Background())

	return
}

// Init 初始化
func (gs *Server) Init(conf *Config) {
	if conf.Name == "" {
		panic("name启动参数不能为空")
	}

	if conf.Version == "" {
		panic("version启动参数不能为空")
	}

	if conf.IP == "" {
		conf.IP = "0.0.0.0"
	}

	if conf.Port <= 0 {
		panic("port启动参数不能为空")
	}

	if conf.BroadcastIP == "" {
		conf.BroadcastIP = conf.IP
	}

	if conf.BroadcastPort <= 0 {
		conf.BroadcastPort = conf.Port
	}

	if conf.Timeout <= 0 {
		panic("timeout启动参数不能为空")
	}

	// 调试模式
	if conf.Debug {
		gs.engine.Use(gin.Logger())
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// 性能监测
	if conf.Pprof {
		pprofGroup := gs.engine.Group("/debug/pprof")

		pprofGroup.GET("/cmdline", func(ctx *gin.Context) {
			pprof.Cmdline(ctx.Writer, ctx.Request)
		})

		pprofGroup.GET("/profile", func(ctx *gin.Context) {
			pprof.Profile(ctx.Writer, ctx.Request)
		})

		pprofGroup.GET("/symbol", func(ctx *gin.Context) {
			pprof.Symbol(ctx.Writer, ctx.Request)
		})

		pprofGroup.GET("/trace", func(ctx *gin.Context) {
			pprof.Trace(ctx.Writer, ctx.Request)
		})

		pprofGroup.GET("/", func(ctx *gin.Context) {
			pprof.Index(ctx.Writer, ctx.Request)
		})

		pprofGroup.GET("/block", func(ctx *gin.Context) {
			pprof.Index(ctx.Writer, ctx.Request)
		})

		pprofGroup.GET("/goroutine", func(ctx *gin.Context) {
			pprof.Index(ctx.Writer, ctx.Request)
		})

		pprofGroup.GET("/heap", func(ctx *gin.Context) {
			pprof.Index(ctx.Writer, ctx.Request)
		})

		pprofGroup.GET("/mutex", func(ctx *gin.Context) {
			pprof.Index(ctx.Writer, ctx.Request)
		})

		pprofGroup.GET("/threadcreate", func(ctx *gin.Context) {
			pprof.Index(ctx.Writer, ctx.Request)
		})
	}

	// 设置 http server
	gs.server = &http.Server{
		Addr:    conf.Addr(),
		Handler: gs.engine,
	}

	gs.Config = conf

	// 加载核心中间件
	gs.engine.Use(core())

	// 加载全局中间件
	gs.Middleware.init()

	// 初始化路由
	gs.Router.init()

	// 加载HTML模板
	// gin.Engine 在创建时，模板尚未初始完毕，需要在这里再进行设置
	// 因 Start 只会调用一次，在 Stop 后应用会直接退出，忽略线程不安全的警告
	if gs.templ != nil {
		gs.engine.SetHTMLTemplate(gs.templ)
	}

	// 加载安全初始化函数
	for _, fn := range gs.initFuncList {
		fn()
	}
}

// AddInit 添加安全初始化函数
func (gs *Server) AddInit(initFunc ...InitFunc) {
	if len(initFunc) > 0 {
		gs.initFuncList = append(gs.initFuncList, initFunc...)
	}
}

// Start 启动服务
func (gs *Server) Start() {

	err := gs.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {

	}
}

// Stop 停止服务
func (gs *Server) Stop() {

	stopCtx, stopCancel := context.WithTimeout(gs.rootCtx, time.Duration(gs.Config.Timeout)*time.Second)

	// FIXME: 不关闭的话，优雅退出时，会导致有连接挂起，总是需要超时退出
	// 目前关闭 keep-alive 状态并未成功
	gs.server.SetKeepAlivesEnabled(false)

	err := gs.server.Shutdown(stopCtx)
	if err != nil {
	} else {
	}

	gs.rootCancel()
	stopCancel()

	// 延时2秒退出，让超时任务 504 响应完成
	time.Sleep(2 * time.Second)
}
