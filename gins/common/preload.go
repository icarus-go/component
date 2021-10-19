package common

import "gorm.io/gorm"

type Preload struct {
	Field string                     `json:"field"`
	Args  func(tx *gorm.DB) *gorm.DB `json:"args"`
}

//NewPreloadByCondition 创建预加载对象并且追加where条件
func NewPreloadByCondition(field string, condition func(tx *gorm.DB) *gorm.DB) *Preload {
	instance := &Preload{
		Field: field,
		Args:  condition,
	}

	return instance
}

//NewDefaultPreload 创建默认预加载
func NewDefaultPreload(field string) *Preload {
	return &Preload{
		Field: field,
	}
}
