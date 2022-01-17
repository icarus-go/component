package etcd

import (
	"github.com/icarus-go/component/etcd/constant"
	"github.com/icarus-go/component/etcd/gin"
	"go.etcd.io/etcd/clientv3"
	"strings"
	"time"
)

//NewGinModeConfig 根据Gin的 gin.Mode 来决定什么环境取什么环境变量
func NewGinModeConfig() (*Config, error) {
	ginMode := gin.NewGinMode()

	value, err := ginMode.GetPath()
	if err != nil {
		return nil, err
	}

	config := new(Config).SetSystem(ginMode).SetConfig(clientv3.Config{
		Endpoints:   strings.Split(value, constant.Comma.Value()),
		DialTimeout: 5 * time.Second,
	})

	config.duration = 5 * time.Second

	return config, nil
}
