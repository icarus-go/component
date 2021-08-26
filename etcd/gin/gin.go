package gin

import (
	"github.com/gin-gonic/gin"
	"os"
	"pmo-test4.yz-intelligence.com/kit/component/etcd/constant"
	"pmo-test4.yz-intelligence.com/kit/component/etcd/errors"
	"regexp"
	"strings"
)

type ginMode struct{}

func NewGinMode() *ginMode {
	return &ginMode{}
}

//GetMode 获取环境变量
func (m *ginMode) GetMode() string {
	return gin.Mode()
}

//GetEnvVarName 获取环境变量名称
func (m *ginMode) GetEnvVarName() (environment constant.Environment, err error) {
	mode := m.GetMode()

	environment = constant.Develop // 默认为develop , 如果没有拿到相应的mode, 会报错

	if mode == gin.TestMode {
		environment = constant.Test // 测试环境
	} else if mode == gin.ReleaseMode {
		environment = constant.Release // 正式环境
	} else if mode == gin.DebugMode {
		environment = constant.Develop // 开发环境
	} else {
		err = &errors.EnvironmentError{Err: errors.Empty, VariableName: environment.Value()} // 客户端可以使用errors.Is来判断是否是错误的变量名称设置
		return
	}
	return
}

//GetPath 根据系列变量名称获取相应的变量值
func (m *ginMode) GetPath() (string, error) {
	name, err := m.GetEnvVarName()
	if err != nil {
		return "", err
	}

	value := os.Getenv(name.Value())
	if value == "" {
		return "", &errors.EnvironmentError{Err: errors.ValueEmpty, VariableName: name.Value()}
	}

	address := strings.Split(value, constant.Comma.Value()) // 集群地址集,可单机
	for _, addr := range address {
		match, err := regexp.Match(constant.AddrRegexp, []byte(addr))
		if err != nil {
			return "", &errors.EnvironmentError{Err: errors.Invalid, VariableName: name.Value(), VariableValue: value, AppendErr: err}
		}

		if !match {
			return "", &errors.EnvironmentError{Err: errors.Invalid, VariableName: name.Value(), VariableValue: value}
		}
	} // 检测值是否符合要求
	return value, nil
}

func (m *ginMode) FileType() string {
	return "yaml"
}
