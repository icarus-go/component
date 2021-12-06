package db

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"pmo-test4.yz-intelligence.com/kit/data/params"
)

type CURL struct {
	instance *gorm.DB
	model    schema.Tabler
}

// SetDBInstance
//  Author: Kevin·CC
//  Description: 设置数据库实例对象
//  Param instance
func (c *CURL) SetDBInstance(instance *gorm.DB) *CURL {
	c.instance = instance
	return c
}

// SetModel
//  Author: Kevin·CC
//  Description: 设置基础model
//  Param md
//  Return *CURL
func (c *CURL) SetModel(md schema.Tabler) *CURL {
	c.model = md
	return c
}

// Delete
//  Author: Kevin·CC
//  Description: 默认删除
//  Param info ID
//  Param scopes 其他条件
//  Param idColumnName id字段名称
//  Return error 错误信息
func (c *CURL) Delete(info params.IWhere, scopes func(*gorm.DB) *gorm.DB, idColumnName ...string) error {
	base := c.instance.Scopes(info.Scopes(idColumnName...))
	if scopes != nil {
		base.Scopes(scopes)
	}
	return base.Delete(c.model).Error
}

// First
//  Author: Kevin·CC
//  Description: 获取单条记录
//  Param info 参数
//  Param value 返回值
//  Param scopes 范围
//  Param idColumnName 主键字段名称
//  Return error 错误信息
func (c *CURL) First(info params.IWhere, value interface{}, scopes func(*gorm.DB) *gorm.DB, idColumnName ...string) error {
	base := c.instance.Scopes(info.Scopes(idColumnName...))
	if scopes != nil {
		base.Scopes(scopes)
	}
	return base.Model(c.model).First(&value).Error
}

// FirstOne
//  Author: Kevin·CC
//  Description: 获取单条记录,如果返回不存在,不报错
//  Param info 条件
//  Param value 值映射对象
//  Param scopes 除ID外其他筛选
//  Param idColumnName ID字段名称
//  Return error 错误信息
func (c *CURL) FirstOne(info params.IWhere, value interface{}, scopes func(*gorm.DB) *gorm.DB, idColumnName ...string) error {
	err := c.First(info, value, scopes, idColumnName...)
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
func (c *CURL) FindOne(info params.IWhere, value interface{}, scopes func(*gorm.DB) *gorm.DB, idColumnName ...string) error {
	base := c.instance.Scopes(info.Scopes(idColumnName...))
	if scopes != nil {
		base.Scopes(scopes)
	}
	return base.Model(c.model).Limit(1).Find(&value).Error
}

// Find
//  Author: Kevin·CC
//  Description: 根据IDs进行搜索
//  Param info 主键筛选
//  Param value 映射值
//  Param scopes 其他筛选
//  Param idColumnName 主键ID名称
//  Return error 错误信息
func (c *CURL) Find(info params.IWhere, value interface{}, scopes func(*gorm.DB) *gorm.DB, idColumnName ...string) error {
	tx := c.instance.Scopes(info.Scopes(idColumnName...))
	if scopes != nil {
		tx.Scopes(scopes)
	}
	return tx.Find(&value).Error
}

// Update
//  Author: Kevin·CC
//  Description: 更新
//  Param info ID更新
//  Param value 值
//  Param scopes 更多的筛选条件
//  Param idColumnName ID字段名称
//  Return error 错误信息
func (c *CURL) Update(info params.IWhere, value interface{}, scopes func(*gorm.DB) *gorm.DB, idColumnName ...string) error {
	base := c.instance.Scopes(info.Scopes(idColumnName...))
	if scopes != nil {
		base.Scopes(scopes)
	}
	return base.Model(c.model).Updates(&value).Error
}
