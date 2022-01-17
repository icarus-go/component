package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/icarus-go/component/gins"
)

// Engine 服务对象
var Engine *gin.Engine

// Init 初始化gins Engine
func Init(conf *gins.Config) {
	gins.Instance.Init(conf)

	Engine = gins.Engine()
}
