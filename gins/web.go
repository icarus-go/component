package gins

import (
	"net/http"
)

type web struct {
	ctx *Context
}

// Render 立即渲染Web
func (w *web) Render(name string, obj interface{}) {
	w.ctx.Context.HTML(http.StatusOK, name, obj)
}
