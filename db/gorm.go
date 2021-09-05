package db

import (
	"database/sql"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"pmo-test4.yz-intelligence.com/kit/component/db/config"
	"pmo-test4.yz-intelligence.com/kit/component/db/log"
)

type Gorm struct {
	DB  *gorm.DB
	SQL *sql.DB

	config.Params

	AutoMigrateTables []AutoMigrateTable
}

func New(config config.Params, autoMigrateTable ...AutoMigrateTable) *Gorm {
	instance := new(Gorm)

	instance.Params = config

	instance.AutoMigrateTables = autoMigrateTable

	return instance
}

// Initialize gorm连接mysql数据库
// Author SliverHorn
func (m *Gorm) Initialize() error {
	var err error

	m.DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		DisableDatetimePrecision:  true,    // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,    // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,    // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}), log.GenerateConfig(m.LogMode, m.LogZap))

	if err != nil {
		zap.L().Error(`Gorm连接MySQL异常!`, zap.Error(err))
		os.Exit(0)
		return err
	}

	if m.AutoMigrate {
		if m.AutoMigrateTables == nil {
			return fmt.Errorf("AutoMigrateTable对象为空,请实现Register()方法")
		}
		for _, table := range m.AutoMigrateTables {
			table.Register(m.DB)
		}
	}

	if m.SQL, err = m.DB.DB(); err != nil {
		zap.L().Error(`DatabaseSql对象获取异常!`, zap.Error(err))
		return err
	}

	m.SQL.SetMaxIdleConns(m.GetMaxIdleConns())
	m.SQL.SetMaxOpenConns(m.GetMaxOpenConns())
	return nil
}
