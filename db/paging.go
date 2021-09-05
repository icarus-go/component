package db

import (
	"gorm.io/gorm"
	"pmo-test4.yz-intelligence.com/kit/component/gins/request"
)

func Paginate(paging *request.Paging) func(db *gorm.DB) *gorm.DB {
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
