package db

import "gorm.io/gorm"

type AutoMigrateTable interface {
	Register(*gorm.DB)
}
