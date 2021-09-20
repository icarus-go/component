package config

type Params struct {
	Path                     string `mapstructure:"path" json:"path" yaml:"path"`
	Config                   string `mapstructure:"config" json:"config" yaml:"config"`
	Dbname                   string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	Username                 string `mapstructure:"username" json:"username" yaml:"username"`
	Password                 string `mapstructure:"password" json:"password" yaml:"password"`
	MaxIdleConns             int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns             int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	LogMode                  string `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`
	AutoMigrate              bool   `mapstructure:"AutoMigrate" json:"autoMigrate" yaml:"AutoMigrate"`
	LogZap                   bool   `mapstructure:"log-zap" json:"logZap" yaml:"log-zap"`
	AllowGlobalUpdate        bool   `mapstructure:"allow-global" json:"allowGlobalUpdate" yaml:"allow-global-update"`
	DisableAutomaticPing     bool   `mapstructure:"disable-automatic-ping" json:"disableAutomaticPing" yaml:"disable-automatic-ping"`
	DisableNestedTransaction bool   `mapstructure:"disable-nested-transaction" json:"disableNestedTransaction" yaml:"disable-nested-transaction"`
	DisableDBStartes         bool   `mapstructure:"disable-db-starts" json:"disableDbStartes" yaml:"disable-db-starts"`
	NamingStrategy
}

type NamingStrategy struct {
	TablePrefix   string `json:"tablePrefix" yaml:"table-prefix"`
	SingularTable bool   `json:"singularTable" yaml:"singular-table"`
	NoLowerCase   bool   `json:"noLowerCase" yaml:"no-lower-case"`
}

func (m *Params) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
}

func (m *Params) GetMaxIdleConns() int {
	return m.MaxIdleConns
}

func (m *Params) GetMaxOpenConns() int {
	return m.MaxOpenConns
}
