// Package db
// Description:
// Author: Kevin · Cai
// Created: 2022/3/4 17:25:54
package db

import (
	"github.com/icarus-go/component/db/config"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

type Resolver struct {
	dbresolver.DBResolver
	masters          []*Gorm
	slaves           []*Gorm
	dbResolverPolicy dbresolver.Policy
}

// NewResolver 初始化读写分离的数据库
// param:
//  separte:
// return:
//  *Resolver:
func newResolver(masterConfigs []config.Params, slaveConfigs []config.Params, fn starter) (*Resolver, error) {
	masters := make([]*Gorm, 0, len(masterConfigs))

	for _, masterConfig := range masterConfigs {
		master, err := New(masterConfig, fn)
		if err != nil {
			return nil, err
		}

		masters = append(masters, master)
	}

	slaves := make([]*Gorm, len(slaveConfigs))
	for _, slaveConfig := range slaveConfigs {
		item, err := New(slaveConfig, fn)
		if err != nil {
			return nil, err
		}

		slaves = append(slaves, item)
	}

	return &Resolver{
		masters: masters,
		slaves:  slaves,
	}, nil
}

// NewResolver 一主多从
//  Param separate 一主多从配置
//  Param fn 方法
//  Return *Resolver: 帮助工具
//  Return error: 错误信息
func NewResolver(separate config.Separate, fn starter) (*Resolver, error) {
	return newResolver([]config.Params{separate.Master}, separate.Slaves, fn)
}

// NewMultiResolver 多主多从
//  Param multipart 多主多从配置
//  Param fn 启动方法
//  Return *Resolver: 帮助工具
//  Return error: 错误信息
func NewMultiResolver(multipart config.Multipart, fn starter) (*Resolver, error) {
	return newResolver(multipart.Masters, multipart.Slaves, fn)
}

// NewDB 初始化DB
//  return *Gorm GORM实例对象
//  return error 错误信息
func (r *Resolver) NewDB() (*Gorm, error) {
	first := r.masters[0] // 首个gorm链接作为返回

	if err := r.register(first.DB); err != nil {
		return nil, err
	}

	return first, nil
}

func (r *Resolver) register(instance *gorm.DB) error {
	return instance.Use(dbresolver.Register(dbresolver.Config{
		Sources:  r.dialectic(r.masters...),
		Replicas: r.dialectic(r.slaves...),
		Policy:   r.policy(),
	}))
}

func (*Resolver) dialectic(dialectic ...*Gorm) []gorm.Dialector {
	dialectics := make([]gorm.Dialector, 0, len(dialectic))
	for _, g := range dialectic {
		dialectics = append(dialectics, g.DB.Dialector)
	}
	return dialectics
}

// SetPolicy 支持自定义负载均衡策略
//  dbResolverPolicy: 负载均衡策略
func (r *Resolver) SetPolicy(dbResolverPolicy dbresolver.Policy) *Resolver {
	r.dbResolverPolicy = dbResolverPolicy
	return r
}

func (r *Resolver) policy() dbresolver.Policy {
	if r.dbResolverPolicy != nil {
		return r.dbResolverPolicy
	}
	return dbresolver.RandomPolicy{}

}
