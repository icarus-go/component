package config

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
