package config

type Params struct {
	Path         string `mapstructure:"path" json:"path" yaml:"path"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	Dbname       string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	LogMode      string `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`
	AutoMigrate  bool   `mapstructure:"AutoMigrate" json:"autoMigrate" yaml:"AutoMigrate"`
	LogZap       bool   `mapstructure:"log-zap" json:"logZap" yaml:"log-zap"`
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
