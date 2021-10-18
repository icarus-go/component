package repo

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"go.uber.org/zap"
	"strconv"
)

type Boolean struct {
	Result *bool
}

// Int 布尔对应数值
//  Author:  Kevin·CC
func (b Boolean) Int() int {
	isOK := 0
	if b.Result != nil && *b.Result {
		isOK = 1
	}
	return isOK
}

// NotNil 是否为空
//  Author:  Kevin·CC
func (b *Boolean) NotNil() bool {
	return b.Result != nil
}

// Scan 扫描
// Author SliverHorn
func (b *Boolean) Scan(value interface{}) error {
	nullBool := sql.NullBool{}
	if err := nullBool.Scan(value); err != nil {

		nullInt := sql.NullInt32{}

		if err := nullInt.Scan(value); err != nil {
			zap.L().Error("Boolean To Database 时间转换失败!", zap.Any("datetime", value))
			return err
		}

		if nullInt.Int32 != 0 {
			result := true
			b.Result = &result
			return nil
		}
		result := false
		b.Result = &result
	}
	b.Result = &nullBool.Bool
	return nil
}

// Value 值
// Author SliverHorn
func (b *Boolean) Value() (driver.Value, error) {
	return driver.Value(strconv.FormatBool(*b.Result)), nil
}

// MarshalJSON 序列化
// Author SliverHorn
func (b *Boolean) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, strconv.FormatBool(*b.Result))), nil
}

// UnmarshalJSON 反序列化
// Author SliverHorn
func (b *Boolean) UnmarshalJSON(bytes []byte) error {
	boolValue, err := strconv.ParseBool(string(bytes))
	if err != nil {
		zap.L().Error("boolean parse fail!", zap.String("value", string(bytes)))
		return err
	}
	b.Result = &boolValue

	return nil
}

// GormDataType gorm 定义数据库字段类型
// Author SliverHorn
func (b *Boolean) GormDataType() string {
	return "tinyint"
}
