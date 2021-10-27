package repo

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"strconv"
)

type Boolean bool

// Int 布尔对应数值
//  Author:  Kevin·CC
func (b Boolean) Int() int {
	if b {
		return 1
	}
	return 0
}

// Scan 扫描
// Author Kevin·CC
func (b *Boolean) Scan(value interface{}) error {
	nullBool := sql.NullBool{}
	err := nullBool.Scan(value)
	*b = Boolean(nullBool.Bool)
	return err
}

// Value 值
// Author Kevin·CC
func (b *Boolean) Value() (driver.Value, error) {
	return driver.Value(strconv.FormatBool(bool(*b))), nil
}

// MarshalJSON 序列化
// Author Kevin·CC
func (b *Boolean) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%v`, *b)), nil
}

// UnmarshalJSON 反序列化
// Author Kevin·CC
func (b *Boolean) UnmarshalJSON(bytes []byte) error {
	boolValue, err := strconv.ParseBool(string(bytes))
	*b = Boolean(boolValue)
	return err
}

// GormDataType gorm 定义数据库字段类型
// Author Kevin·CC
func (b *Boolean) GormDataType() string {
	return "tinyint"
}
