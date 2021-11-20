package db

import "gorm.io/gorm"

//AutoMigrateTable
//  Author: Kevin·CC
//  Description: 自动注册接口,每个模块需要注册哪些内容
type AutoMigrateTable interface {
	Register(*gorm.DB)
}
