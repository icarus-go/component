// Package gins .
package gins

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"html/template"
	"net/http"
	"os"

	"syscall"

	"os/signal"

	"github.com/gin-gonic/gin"
)

var (
	// Instance gins实例
	Instance   *Server
	signalChan chan os.Signal
)

func init() {
	var err error
	if Instance, err = New(); err != nil {
		err = errors.New("Gin Server 创建失败：" + err.Error())
		panic(err)
	}

	signalChan = make(chan os.Signal)
}

// Run 启动GinServer
func Run(conf *Config) {

	// 初始化服务配置
	Instance.Init(conf)

	if !conf.IsDisableSignal {
		//使用docker stop 命令去关闭Container时，该命令会发送SIGTERM 命令到Container主进程，让主进程处理该信号，关闭Container，如果在10s内，未关闭容器，Docker Damon会发送SIGKILL 信号将Container关闭。
		signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT)
	}

	baseURL := fmt.Sprintf("%s:%d", conf.IP, conf.Port)

	zap.L().Info(fmt.Sprintf(`
当前应用: %s
当前版本: %s
默认自动化文档地址:http://%s/swagger/index.html
默认前端文件运行地址:http://%s`, conf.Name, conf.Version, baseURL, baseURL))

	go func() {
		Instance.Start()
	}()

	<-signalChan

	if !conf.IsDisableSignal {
		signal.Stop(signalChan)
	}

	Instance.Stop()

	if conf.IsDisableSignal {
		signalChan <- syscall.SIGINT
	}
}

// Stop 停止GinServer
func Stop() {
	signalChan <- syscall.SIGINT

	if Instance.Config.IsDisableSignal {
		<-signalChan
	}
}

// AddInit 添加安全初始化函数
func AddInit(initFuncs ...InitFunc) {
	Instance.AddInit(initFuncs...)
}

// AddTemplate 添加模板
func AddTemplate(name string, content string) {
	if Instance.templ == nil {
		Instance.templ = template.Must(template.New(name).Parse(content))
	} else {
		Instance.templ = template.Must(Instance.templ.New(name).Parse(content))
	}
}

// RouterGroup 路由对象
func RouterGroup(relativePath string, handlers ...HandlerFunc) *Router {
	return Instance.Router.Group(relativePath, handlers...)
}

// Use 全局中间件
func Use(middlewares ...HandlerFunc) {
	Instance.Middleware.Use(middlewares...)
}

// POST 请求
func POST(relativePath string, handlers ...HandlerFunc) *Router {
	return Instance.Router.POST(relativePath, handlers...)
}

//AddPlugins 添加插件
func AddPlugins(group *Router, plugins ...Plugin) {
	for _, plugin := range plugins {
		plugin.Register(group.Group(plugin.Path()))
	}
}

// GET 请求
func GET(relativePath string, handlers ...HandlerFunc) *Router {
	return Instance.Router.GET(relativePath, handlers...)
}

// StaticFile 静态文件服务
func StaticFile(relativePath, filepath string) *Router {
	return Instance.Router.StaticFile(relativePath, filepath)
}

// Static 静态文件目录服务
func Static(relativePath, root string) *Router {
	return Instance.Router.Static(relativePath, root)
}

// StaticFS 静态资源服务
func StaticFS(relativePath string, fs http.FileSystem) *Router {
	return Instance.Router.StaticFS(relativePath, fs)
}

// On404 404自定义处理
func On404(handler HandlerFunc) {
	Instance.on404 = handler
}

// On500 500自定义处理
func On500(handler HandlerFunc) {
	Instance.on500 = handler
}

// Engine gin.Engine对象
// 本框架对Router进行了二次封装，调用此对象时请注意路由冲突问题
func Engine() *gin.Engine {
	return Instance.engine
}
