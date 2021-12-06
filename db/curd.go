package db

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"pmo-test4.yz-intelligence.com/kit/data/params"
)

type CURL struct {
	Gorm
	Model schema.Tabler
}

// Delete
//  Author: Kevin·CC
//  Description: 默认删除
//  Param info ID
//  Param scopes 其他条件
//  Param idColumnName id字段名称
//  Return error 错误信息
func (m *CURL) Delete(info params.IWhere, scopes func(*gorm.DB) *gorm.DB, idColumnName ...string) error {
	base := m.DB.Scopes(info.Scopes(idColumnName...))
	if scopes != nil {
		base.Scopes(scopes)
	}
	return base.Delete(m.Model).Error
}

// First
//  Author: Kevin·CC
//  Description: 获取单条记录
//  Param info 参数
//  Param value 返回值
//  Param scopes 范围
//  Param idColumnName 主键字段名称
//  Return error 错误信息
func (m *CURL) First(info params.IWhere, value interface{}, scopes func(*gorm.DB) *gorm.DB, idColumnName ...string) error {
	base := m.DB.Scopes(info.Scopes(idColumnName...))
	if scopes != nil {
		base.Scopes(scopes)
	}
	return base.First(&value).Error
}

// FirstOne
//  Author: Kevin·CC
//  Description: 获取单条记录,如果返回不存在,不报错
//  Param info 条件
//  Param value 值映射对象
//  Param scopes 除ID外其他筛选
//  Param idColumnName ID字段名称
//  Return error 错误信息
func (m *CURL) FirstOne(info params.IWhere, value interface{}, scopes func(*gorm.DB) *gorm.DB, idColumnName ...string) error {
	err := m.First(info, value, scopes, idColumnName...)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return err
}

// FindOne
//  Author: Kevin·CC
//  Description: 获取单挑
//  Param info
//  Param value
//  Param scopes
//  Param idColumnName
//  Return error
func (m *CURL) FindOne(info params.IWhere, value interface{}, scopes func(*gorm.DB) *gorm.DB, idColumnName ...string) error {
	base := m.DB.Scopes(info.Scopes(idColumnName...))
	if scopes != nil {
		base.Scopes(scopes)
	}
	return base.Limit(1).Find(&value).Error
}

// Update
//  Author: Kevin·CC
//  Description: 更新
//  Param info ID更新
//  Param value 值
//  Param scopes 更多的筛选条件
//  Param idColumnName ID字段名称
//  Return error 错误信息
func (m *CURL) Update(info params.IWhere, value interface{}, scopes func(*gorm.DB) *gorm.DB, idColumnName ...string) error {
	base := m.DB.Scopes(info.Scopes(idColumnName...))
	if scopes != nil {
		base.Scopes(scopes)
	}
	return base.Updates(&value).Error
}
