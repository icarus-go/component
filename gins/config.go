package gins

import "fmt"

// Config 服务器配置
type Config struct {
	Name            string // 服务名称，必填
	Version         string // 服务版本，必填
	Host            string // 域名主机
	IP              string // 运行地址，必填
	Port            int    // 运行端口，必填
	Timeout         int    // 优雅退出时的超时机制
	Debug           string // 是否开启调试
	Pprof           bool   // 是否监控性能
	Cors            bool   // 是否支持跨域处理
	IsDisableSignal bool   // 是否关闭 signal 信号监听退出，默认：false，设置为 true 时，需主动调用 gins.Stop() 来触发优雅退出
}

// Addr 运行地址
func (conf *Config) Addr() string {
	return fmt.Sprintf("%s:%d", conf.IP, conf.Port)
}
