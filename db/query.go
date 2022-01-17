package db

import (
	"errors"
	"github.com/icarus-go/data/params"
	"gorm.io/gorm"
)

// MustPaginate
//  Author: Kevin·CC
//  Description: 必须有分页参数
//  Param paging 分页
//  Return func(db *gorm.DB) *gorm.DB 分页 limit , offset
//  Return error 错误信息
func MustPaginate(paging *params.Paging) (func(db *gorm.DB) *gorm.DB, error) {
	if paging == nil {
		return nil, errors.New("分页参数为空")
	}
	return Paginate(paging), nil
}

// Paginate
//  Author: Kevin·CC
//  Description: 分页非必须 适合通用方法/基础方法
//  Param paging 分页
//  Return func(db *gorm.DB) *gorm.DB 分页返回
func Paginate(paging *params.Paging) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if paging != nil {
			switch {
			case paging.PageSize > 10000:
				paging.PageSize = 100
			case paging.PageSize < 0:
				paging.PageSize = 20
			}
			offset := paging.PageSize * (paging.Page - 1)
			return db.Offset(offset).Limit(paging.PageSize)
		}
		return db
	}
}

//Order 排序方法
//  orders 排序对象
func Order(orders ...*params.Order) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for _, order := range orders {
			db.Order(order.Join() + " " + order.Sort())
		}
		return db
	}
}

// Preload
//  Author: Kevin·CC
//  Description: 预加载
//  Param preload 预加载对象
//  Return func(db *gorm.DB) *gorm.DB 限制
func Preload(preload ...*params.Preload) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for _, item := range preload {
			if item.Args == nil {
				db.Preload(item.Field)
				continue
			}

			db.Preload(item.Field, item.Args)
		}
		return db
	}
}

// Restrict
//  Author: Kevin·CC
//  Description: 限制预加载, 必须继承接口的实现
//  Param restrict 限制参数
//  Return func(tx *gorm.DB) *gorm.DB 限制
func Restrict(restrict ...params.RestrictPreload) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		for _, item := range restrict {

			if item.IRestrict == nil {
				tx.Preload(item.Object)
				continue
			}

			tx.Preload(item.Object, item.Where())
		}
		return tx
	}
}
