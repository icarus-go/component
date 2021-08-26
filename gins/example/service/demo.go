package service

import (
	"pmo-test4.yz-intelligence.com/kit/component/gins"
	"pmo-test4.yz-intelligence.com/kit/component/gins/example/db"
	"pmo-test4.yz-intelligence.com/kit/component/gins/example/errors"
	"pmo-test4.yz-intelligence.com/kit/component/gins/example/model"
)

// Demo 用户业务逻辑
var Demo demo

type demo struct{}

// Get 获取
func (*demo) Get(ctx *gins.Context, id string) (md *model.Demo, err error) {
	// 如果需要登录
	// _, err = CheckLogin(ctx)
	// if err != nil {
	// 	return
	// }

	if id == "" {
		err = errors.New("ID不能为空")
		return
	}

	if err = db.Instance.Create(md).Error; err != nil {
		return nil, err
	}

	return md, nil
}

// Add 插入数据
func (*demo) Add(ctx *gins.Context, md *model.Demo) (err error) {
	// 如果需要登录
	// _, err = CheckLogin(ctx)
	// if err != nil {
	// 	return
	// }

	// 检查名称是否已被占用s
	err = Demo.CheckName(ctx, md.Name, "")
	if err != nil {
		return
	}

	return
}

// Update 更新
func (*demo) Update(ctx *gins.Context, md *model.Demo) (affected int64, err error) {
	// 如果需要登录
	// _, err = CheckLogin(ctx)
	// if err != nil {
	// 	return
	// }

	return
}

// Delete 删除
func (*demo) Delete(ctx *gins.Context, id string) (affected int64, err error) {

	return
}

// List 列表
func (*demo) List(ctx *gins.Context, qry interface{}) (ml interface{}, err error) {

	return
}

// CheckName 检查名称
func (*demo) CheckName(ctx *gins.Context, name, id string) (err error) {
	//
	//md := &model.Demo{}
	//da, err := orm.NewDA(md)
	//if err != nil {
	//	err = errors.Except(err)
	//	return
	//}
	//defer da.Close()
	//
	//has, err := da.Where("name=?", name).Get(md)
	//
	//if err != nil {
	//	err = errors.Except(err)
	//	return
	//}
	//
	//if has {
	//
	//	if md.ID == id {
	//		return
	//	}
	//
	//	err = errors.LogicMask(code.DEMO_DATA_ADD_Had) // 使用了自定义错误信息，但仍用90000码输出。如果code想用自定义码，使用errors.Logic(code.DEMO_DATA_ADD_HAD)
	//	return
	//}

	return
}
