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

//NamingStrategy
//  Author: Kevin·CC
//  Description: 命名池
type NamingStrategy struct {
	TablePrefix   string `json:"tablePrefix" yaml:"table-prefix"`     //  TablePrefix 表前缀
	SingularTable bool   `json:"singularTable" yaml:"singular-table"` //  SingularTable 是否名称为单数
	NoLowerCase   bool   `json:"noLowerCase" yaml:"no-lower-case"`    //  NoLowerCase 是否非小写
}

//Dsn
//  Author: Kevin·CC
//  Description: 数据库连接字符串
//  Return string
func (m *Params) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
}

//GetMaxIdleConns
//  Author: Kevin·CC
//  Description: 最大等待连接数
//  Return int
func (m *Params) GetMaxIdleConns() int {
	return m.MaxIdleConns
}

//GetMaxOpenConns
//  Author: Kevin·CC
//  Description: 最大连接数
//  Return int
func (m *Params) GetMaxOpenConns() int {
	return m.MaxOpenConns
}
