package service

import (
	"time"

	"pmo-test4.yz-intelligence.com/kit/component/gins"
)

// Info 服务信息
type Info struct {
	bill map[string]interface{} `json:"bill"`
}

// NewInfo 实例
func NewInfo() *Info {
	info := &Info{}
	info.bill = make(map[string]interface{})
	return info
}

// Bill 清单
func (info *Info) Bill() (map[string]interface{}, error) {
	return info.bill, nil
}

// App 应用信息
func (info *Info) App(app gins.Config) *Info {
	info.bill["app"] = app
	return info
}

func (info *Info) Version(version string) *Info {
	info.bill["version"] = version
	return info
}

func (info *Info) Env(env string) *Info {
	info.bill["env"] = env
	return info
}

func (info *Info) ServerTime() *Info {
	info.bill["serverTime"] = time.Now().Unix()
	return info
}

//Configuration 配置信息
func (info *Info) Configuration(configuration interface{}) *Info {
	info.bill["configuration"] = configuration
	return info
}
