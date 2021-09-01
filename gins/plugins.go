package gins

import (
	"fmt"
	"net/http"
	"runtime"
	"sync"

	"github.com/gin-gonic/gin"
	"pmo-test4.yz-intelligence.com/kit/component/apiconstant"
)

var ctxPool sync.Pool

func init() {
	ctxPool.New = func() interface{} {
		return &Context{}
	}
}

// recovery panic恢复，初始化Context
func recovery() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		ctx := ctxPool.New().(*Context)
		ctx.reset(ginCtx)
		ctx.Set("*Context", ctx)

		go func() {
			defer func() {
				//异常捕获处理
				if e := recover(); e != nil {
					stack := fmt.Sprintf("System Panic: %v", e)
					for i := 1; ; i++ {
						_, file, line, ok := runtime.Caller(i)
						if !ok {
							break
						} else {
							stack += "\n"
						}
						stack += fmt.Sprintf("%v:%v", file, line)
					}

					// 500 处理
					if Instance.on500 != nil {
						ctx.setPanic(stack)
						Instance.on500(ctx)
					} else {
						// 默认异常响应
						if ctx.IsAPI() {
							ctx.API.SetError(NewAPIErrorWithLog("系统异常", stack))
							ctx.API.Render()
							ctx.Abort()
						} else {
							ctx.AbortWithStatus(http.StatusInternalServerError)
						}
					}
				}

				close(ctx.doneChan)
			}()
			ctx.Next()
		}()

		<-ctx.doneChan

		if ctx.IsAborted() {
			ctxPool.Put(ctx)
			return
		}

		status := ctx.Writer.Status()

		// 路由匹配到的情况下，status默认 200
		if status != http.StatusNotFound && ctx.IsAPI() {
			if ctx.API.result.Code == apiconstant.RESPONSE_UNKNOW {
				ctx.API.result.Msg = "API空响应"
			}

			ctx.API.Render()
		}

		if status == http.StatusNotFound {
			// 404 处理
			if Instance.on404 != nil {
				Instance.on404(ctx)
			} else {
				ctx.AbortWithStatus(404)
			}
		}

		ctxPool.Put(ctx)
	}
}

//cors 处理跨域请求,支持options访问
func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		origin := c.Request.Header.Get("Origin")

		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

//Plugin 方便注册第三方接口
type Plugin interface {
	// Register 注册路由
	Register(group *Router)

	// Path 用户返回注册路由
	Path() string
}
