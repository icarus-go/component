package mock

import (
	"github.com/gin-gonic/gin"
	"pmo-test4.yz-intelligence.com/kit/component/gins"
)

// Engine 服务对象
var Engine *gin.Engine

// Init 初始化gins Engine
func Init(conf *gins.Config) {
	gins.Instance.Init(conf)

	Engine = gins.Engine()
}
