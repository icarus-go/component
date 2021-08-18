package api

import (
	"time"

	"pmo-test4.yz-intelligence.com/kit/component/gins/example/model"
	"pmo-test4.yz-intelligence.com/kit/component/gins/example/service"

	"pmo-test4.yz-intelligence.com/kit/component/gins"
)

// Demo 测试模块
var Demo demo

type demo struct{}

func (*demo) Add(ctx *gins.Context) {
	md := new(model.Demo)
	err := ctx.ShouldBind(ctx.Request.Body) // application/json；如果是x-www-form-urlencoded用ctx.PostForm("id")
	if err != nil {
		ctx.API.SetError(err)
		return
	}

	err = service.Demo.Add(ctx, md)
	if err != nil {
		ctx.API.SetError(err)
		return
	}

	ctx.API.SetData(md)
	return
}

func (*demo) Update(ctx *gins.Context) {
	md := new(model.Demo)
	err := ctx.ShouldBind(ctx.Request.Body) // application/json；如果是x-www-form-urlencoded用ctx.PostForm("id")
	if err != nil {
		ctx.API.SetError(err)
		return
	}

	var affected int64
	affected, err = service.Demo.Update(ctx, md)
	if err != nil {
		ctx.API.SetError(err)
		return
	}

	if affected <= 0 {
		ctx.API.SetMsg("更新失败，没有ID为 " + md.ID + " 的记录")
		return
	}

	setResponseOK(ctx)
}

func (*demo) Get(ctx *gins.Context) {
	query := ctx.Request.URL.Query()
	id := query.Get("id")

	md, err := service.Demo.Get(ctx, id)

	if err != nil {
		ctx.API.SetError(err)
		return
	}

	ctx.API.SetData(md)
}

// List 列表
func (*demo) List(ctx *gins.Context) {
	md := new(model.Demo)
	err := ctx.ShouldBind(ctx.Request.Body) // application/json；如果是x-www-form-urlencoded用ctx.PostForm("id")
	if err != nil {
		ctx.API.SetError(err)
		return
	}

	ml, err := service.Demo.List(ctx, md)
	if err != nil {
		ctx.API.SetError(err)
		return
	}

	ctx.API.SetData(ml)
}

func (*demo) Post(ctx *gins.Context) {
	ctx.API.SetMsg("contentType：" + ctx.Request.Header.Get("Content-Type"))
}

func (*demo) Timeout(ctx *gins.Context) {
	//测试超时
	time.Sleep(30 * time.Second)
}

func (*demo) Panic(ctx *gins.Context) {
	//测试超时
	panic("测试panic")
}
