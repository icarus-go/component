package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"pmo-test4.yz-intelligence.com/kit/component/db/config"
	thisLog "pmo-test4.yz-intelligence.com/kit/component/db/log"
)

type Gorm struct {
	DB  *gorm.DB
	SQL *sql.DB

	config.Params

	gormConfig *gorm.Config

	AutoMigrateTables []AutoMigrateTable
}

//starter 启动器配置项
type starter func(instance *Gorm)

//DefaultNew 默认仅仅开启（DDL表名规范、自动注册表、日志）
func DefaultNew(config config.Params) (*Gorm, error) {
	return New(config, func(instance *Gorm) {
		instance.
			SetDDLRule().
			SetAutoMigrateTables().
			SetLogger()
	})
}

func New(config config.Params, setGormConfig starter) (*Gorm, error) {
	instance := new(Gorm)

	instance.Params = config

	instance.newGormConfig()

	setGormConfig(instance)

	if err := instance.initialize(); err != nil {
		return nil, err
	}

	return instance, nil
}

// initialize gorm连接mysql数据库
func (m *Gorm) initialize() error {
	var err error

	_ = m.newGormConfig()

	m.DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		DisableDatetimePrecision:  true,    // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,    // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,    // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}), m.gormConfig)

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

func (m *Gorm) newGormConfig() *gorm.Config {
	if m.gormConfig == nil {
		m.gormConfig = &gorm.Config{}
	}
	return m.gormConfig
}

//SetDDLRule 设置DB生成DDL语句配置
func (m *Gorm) SetDDLRule() *Gorm {
	m.gormConfig.NamingStrategy = schema.NamingStrategy{
		TablePrefix:   m.Params.TablePrefix,   // 表名前缀
		SingularTable: m.Params.SingularTable, // 是否单数表名
		NoLowerCase:   m.NoLowerCase,          // 是否不需要 转为 带下划线的表名
	}

	m.gormConfig.DisableForeignKeyConstraintWhenMigrating = true // 默认禁用外检关联和约束
	return m
}

// SetAllowGlobalUpdate 设置是否允许全局更新，默认不允许
func (m *Gorm) SetAllowGlobalUpdate() *Gorm {
	m.gormConfig.AllowGlobalUpdate = m.Params.AllowGlobalUpdate
	return m
}

func (m *Gorm) SetLogger() *Gorm {
	thisLog.Set(m.gormConfig, m.Params)
	return m
}

// SetAutoMigrateTables 设置自动注册表格
func (m *Gorm) SetAutoMigrateTables(autoMigrateTable ...AutoMigrateTable) *Gorm {
	m.AutoMigrateTables = autoMigrateTable
	return m
}

// SetDisableAutomaticPing 设置是否关闭自动PingDB
func (m *Gorm) SetDisableAutomaticPing() *Gorm {
	m.gormConfig.DisableAutomaticPing = m.Params.DisableAutomaticPing
	return m
}

// SetDisableNestedTransaction 设置是否禁止嵌套事务
func (m *Gorm) SetDisableNestedTransaction() *Gorm {
	m.gormConfig.DisableNestedTransaction = m.Params.DisableNestedTransaction
	return m
}

//SetNowFunc 设置当前时间变更方法
func (m *Gorm) SetNowFunc(nowFunc func() time.Time) *Gorm {
	m.gormConfig.NowFunc = nowFunc
	return m
}
