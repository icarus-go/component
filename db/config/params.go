// Package config
// Description: The database param model
// Author: Kevin · Cai
// Created: 2022/3/4 17:07:50
package config

//Params
//  Author: Kevin·CC
//  Description: 数据库层参数配置
type Params struct {
	Path                     string `mapstructure:"path" json:"path" yaml:"path"`                                                                 // 路径
	Config                   string `mapstructure:"config" json:"config" yaml:"config"`                                                           // 配置
	Dbname                   string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`                                                         // 数据库名称
	Username                 string `mapstructure:"username" json:"username" yaml:"username"`                                                     //  Username 用户名称
	Password                 string `mapstructure:"password" json:"password" yaml:"password"`                                                     //  Password 密码
	MaxIdleConns             int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`                                     //  MaxIdleConns 最大等待连接数
	MaxOpenConns             int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`                                     //  MaxOpenConns 最大连接数
	LogMode                  string `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`                                                      //  LogMode 日志模式
	AutoMigrate              bool   `mapstructure:"AutoMigrate" json:"autoMigrate" yaml:"AutoMigrate"`                                            //  AutoMigrate 是否自动注册表
	LogZap                   bool   `mapstructure:"log-zap" json:"logZap" yaml:"log-zap"`                                                         //  LogZap
	AllowGlobalUpdate        bool   `mapstructure:"allow-global" json:"allowGlobalUpdate" yaml:"allow-global-update"`                             //  AllowGlobalUpdate 是否允许不带where进行更新
	DisableAutomaticPing     bool   `mapstructure:"disable-automatic-ping" json:"disableAutomaticPing" yaml:"disable-automatic-ping"`             //  DisableAutomaticPing 是否关闭心跳
	DisableNestedTransaction bool   `mapstructure:"disable-nested-transaction" json:"disableNestedTransaction" yaml:"disable-nested-transaction"` //  DisableNestedTransaction 是否关闭最小单元事务
	DisableDBStartes         bool   `mapstructure:"disable-db-starts" json:"disableDbStartes" yaml:"disable-db-starts"`
	NamingStrategy
}

// Separate 读写分离参数结构体 适用于一主多从
//  Author: Kevin
//  Date: 2022-03-04 17:12:58
type Separate struct {
	Master Params   `yaml:"master" json:"master" mapstructure:"master"` // Master 主库
	Slaves []Params `yaml:"slaves" json:"slaves" mapstructure:"slaves"` // Slaves 从库
}

// Multipart 多主多从
//  Author: Kevin
//  Date: 2022-03-04 17:15:21
type Multipart struct {
	Masters []Params `yaml:"masters" json:"masters" mapstructure:"masters"` // Masters 主库们
	Slaves  []Params `yaml:"slaves" json:"slaves" mapstructure:"slaves"`    // Slaves 从库
}
