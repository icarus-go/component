package etcd

import (
	"github.com/icarus-go/component/etcd/constant"
)

type System interface {
	// GetMode 返回环境, 如果使用Gin则返回gins.GetMode()
	// 如果是通过超
	GetMode() string

	GetEnvVarName() (constant.Environment, error)

	GetPath() (string, error)

	FileType() string
}
