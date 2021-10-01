package db

import (
	"gorm.io/gorm"
	"pmo-test4.yz-intelligence.com/kit/component/gins/common"
)

//Paginate 分页方法
func Paginate(paging *common.Paging) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		switch {
		case paging.PageSize > 10000:
			paging.PageSize = 100
		case paging.PageSize < 0:
			paging.PageSize = 20
		}
		offset := paging.PageSize * (paging.Page - 1)
		return db.Offset(offset).Limit(paging.PageSize)
	}
}

//Order 排序方法
//  orders 排序对象
func Order(orders ...*common.Order) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for _, order := range orders {
			db.Order(order.Join() + " " + order.Sort())
		}
		return db
	}
}

func Cos(cos ...*common.Cos) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		for _, cos := range cos {
			db.Select(cos.Join(), cos.Args)
		}

		return db
	}
}
