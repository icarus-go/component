package gins

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Router 路由
type Router struct {
	*gin.RouterGroup

	tree []func(routerGroup *gin.RouterGroup)
	next []*Router

	action map[string]HandlerFunc
}

// init 初始化路由树
func (r *Router) init() {
	for i, l := 0, len(r.tree); i < l; i++ {
		// 添加路由
		r.tree[i](r.RouterGroup)
	}

	for x, y := 0, len(r.next); x < y; x++ {
		// 设定父级路由，执行子路由初始化
		r.next[x].RouterGroup = r.RouterGroup
		r.next[x].init()
	}
}

// Use 中间件
func (r *Router) Use(middlewares ...HandlerFunc) *Router {
	r.tree = append(r.tree, func(routerGroup *gin.RouterGroup) {
		routerGroup.Use(newGinHandler(middlewares...)...)
	})

	return r
}

// Group 路由分组
func (r *Router) Group(relativePath string, handlers ...HandlerFunc) *Router {
	router := &Router{}

	router.tree = append(router.tree, func(routerGroup *gin.RouterGroup) {
		router.RouterGroup = routerGroup.Group(relativePath, newGinHandler(handlers...)...)
	})

	r.next = append(r.next, router)

	return router
}

// Handle 添加路由方法
func (r *Router) Handle(httpMethod, relativePath string, handlers ...HandlerFunc) *Router {
	r.tree = append(r.tree, func(routerGroup *gin.RouterGroup) {
		routerGroup.Handle(httpMethod, relativePath, newGinHandler(handlers...)...)
	})

	return r
}

// POST 请求
func (r *Router) POST(relativePath string, handlers ...HandlerFunc) *Router {
	r.tree = append(r.tree, func(routerGroup *gin.RouterGroup) {
		routerGroup.POST(relativePath, newGinHandler(handlers...)...)
	})

	return r
}

// GET 请求
func (r *Router) GET(relativePath string, handlers ...HandlerFunc) *Router {
	r.tree = append(r.tree, func(routerGroup *gin.RouterGroup) {
		routerGroup.GET(relativePath, newGinHandler(handlers...)...)
	})

	return r
}

// DELETE 请求
func (r *Router) DELETE(relativePath string, handlers ...HandlerFunc) *Router {
	r.tree = append(r.tree, func(routerGroup *gin.RouterGroup) {
		routerGroup.DELETE(relativePath, newGinHandler(handlers...)...)
	})

	return r
}

// PATCH 请求
func (r *Router) PATCH(relativePath string, handlers ...HandlerFunc) *Router {
	r.tree = append(r.tree, func(routerGroup *gin.RouterGroup) {
		routerGroup.PATCH(relativePath, newGinHandler(handlers...)...)
	})

	return r
}

// PUT 请求
func (r *Router) PUT(relativePath string, handlers ...HandlerFunc) *Router {
	r.tree = append(r.tree, func(routerGroup *gin.RouterGroup) {
		routerGroup.PUT(relativePath, newGinHandler(handlers...)...)
	})

	return r
}

// OPTIONS 请求
func (r *Router) OPTIONS(relativePath string, handlers ...HandlerFunc) *Router {
	r.tree = append(r.tree, func(routerGroup *gin.RouterGroup) {
		routerGroup.OPTIONS(relativePath, newGinHandler(handlers...)...)
	})

	return r
}

// HEAD 请求
func (r *Router) HEAD(relativePath string, handlers ...HandlerFunc) *Router {
	r.tree = append(r.tree, func(routerGroup *gin.RouterGroup) {
		routerGroup.HEAD(relativePath, newGinHandler(handlers...)...)
	})

	return r
}

// Any 任何请求
// GET, POST, PUT, PATCH, HEAD, OPTIONS, DELETE, CONNECT, TRACE.
func (r *Router) Any(relativePath string, handlers ...HandlerFunc) *Router {
	r.tree = append(r.tree, func(routerGroup *gin.RouterGroup) {
		routerGroup.Any(relativePath, newGinHandler(handlers...)...)
	})

	return r
}

// StaticFile 静态文件服务
func (r *Router) StaticFile(relativePath, filepath string) *Router {
	r.tree = append(r.tree, func(routerGroup *gin.RouterGroup) {
		routerGroup.StaticFile(relativePath, filepath)
	})

	return r
}

// Static 静态文件目录服务
func (r *Router) Static(relativePath, root string) *Router {
	r.tree = append(r.tree, func(routerGroup *gin.RouterGroup) {
		routerGroup.Static(relativePath, root)
	})

	return r
}

// StaticFS 静态资源服务
func (r *Router) StaticFS(relativePath string, fs http.FileSystem) *Router {
	r.tree = append(r.tree, func(routerGroup *gin.RouterGroup) {
		routerGroup.StaticFS(relativePath, fs)
	})

	return r
}

// Action Action风格路由中间件
// 在RouterGroup下使用中间件时，必须 路由路径 匹配时才会触发，全局使用时无此问题
func Action(prefixPath ...string) HandlerFunc {
	return func(ctx *Context) {
		query := ctx.Request.URL.Query()

		action := query.Get("action")
		if action == "" {
			// 非Action风格，退回
			ctx.Next()
			return
		}

		actions := strings.Split(action, ".")

		// 增加跳转地址前缀
		urlPath := "/" + strings.Join(actions, "/") + "/"
		if len(prefixPath) == 1 {
			urlPath = prefixPath[0] + urlPath
		}

		// 更新目标
		query.Del("action")
		ctx.Request.URL.Path = urlPath
		ctx.Request.URL.RawQuery = query.Encode()
		Instance.engine.HandleContext(ctx.Context)

		// 命中action风格处理，中止后续行为
		ctx.Abort()
	}
}
