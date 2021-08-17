package gins

import (
	"fmt"
	"net/http"
	"runtime"
	"sync"

	"github.com/gin-gonic/gin"
	"pmo-test4.yz-intelligence.com/kit/apiconstant"
)

var ctxPool sync.Pool

func init() {
	ctxPool.New = func() interface{} {
		return &Context{}
	}
}

// core panic恢复，初始化Context
func core() gin.HandlerFunc {
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

		// FIXME: 存在内存泄露问题
		//监听上下文状态
		// select {
		// case <-ctx.doneChan:
		// 	//正常结束
		// case <-ctx.routerCtx.Done():
		// 	ctx.routerCancel()

		// 	// TODO: 记录更多完善的中断信息
		// 	logger.Warning("请求被强制中断")

		// 	// 退出时请求超 Config.Timeout，被强制取消，响应 504
		// 	ctx.AbortWithStatus(http.StatusGatewayTimeout)
		// }

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
