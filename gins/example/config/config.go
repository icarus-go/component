package config

import (
	"errors"
	"fmt"
	"pmo-test4.yz-intelligence.com/kit/component/db/config"
	"pmo-test4.yz-intelligence.com/kit/component/etcd"
	"pmo-test4.yz-intelligence.com/kit/component/gins/example/arg"
	"pmo-test4.yz-intelligence.com/kit/component/zap"
)

type Config struct {
	MySQL     config.Params `json:"mysql" yaml:"mysql"`
	ZapConfig zap.Config    `json:"zap" yaml:"zap"`
}

var Instance *Config

func Init() error {
	ginModeConfig, err := etcd.NewGinModeConfig()
	if err != nil {
		return err
	}

	ginModeConfig.SetConfigName(arg.Name, ginModeConfig.System.GetMode())

	content, err := ginModeConfig.GetByName()
	if err != nil {
		return err
	}

	if content.Count < 1 {
		return errors.New("获取配置失败")
	}

	if err := ginModeConfig.Unmarshal(content.Kvs[0].Value, &Instance); err != nil {
		return err
	}

	Instance.ZapConfig.Prefix = fmt.Sprintf("[%s]", arg.Name)
	zap.NewZap(Instance.ZapConfig).Initialize() // 根據配置初始化日誌配置

	return nil
}
