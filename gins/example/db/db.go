package db

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
	"pmo-test4.yz-intelligence.com/kit/component/db"
	"pmo-test4.yz-intelligence.com/kit/component/gins/example/config"
	"pmo-test4.yz-intelligence.com/kit/component/gins/example/model"
)

var Instance *gorm.DB

func Init() error {
	m := new(MySQLMigrateTable)

	i := db.New(config.Instance.MySQL, m) // 初始化配置

	if err := i.Initialize(); err != nil {
		return err
	}

	Instance = i.DB

	return nil
}

type MySQLMigrateTable struct{}

func (MySQLMigrateTable) Register(g *gorm.DB) {
	err := g.Debug().AutoMigrate(new(model.Demo))
	if err != nil {
		zap.L().Error(`注册表失败!`, zap.Error(err))
		os.Exit(0)
	}
	zap.L().Info(`注册表成功!`)
}
