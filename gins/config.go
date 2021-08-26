package gins

import (
	"fmt"
)

// Config 服务器配置
type Config struct {
	Name            string `json:"-" yaml:"-"`                             // 服务名称，由调用方指定
	Version         string `json:"-" yaml:"-"`                             // 服务版本，由调用方指定
	Host            string `json:"host" yaml:"host"`                       // 域名主机
	IP              string `json:"ip" yaml:"ip"`                           // 运行地址，必填
	Port            int    `json:"port" yaml:"port"`                       // 运行端口，必填
	Timeout         int    `json:"timeout" yaml:"timeout"`                 // 优雅退出时的超时机制
	Debug           string `json:"debug" yaml:"debug"`                     // 是否开启调试
	Pprof           bool   `json:"pprof" yaml:"pprof"`                     // 是否监控性能
	IsCorsDisable   bool   `json:"isCorsDisable" yaml:"isCorsDisable"`     // 是否支持跨域处理
	IsDisableSignal bool   `json:"isDisableSignal" yaml:"isDisableSignal"` // 是否关闭 signal 信号监听退出，默认：false，设置为 true 时，需主动调用 gins.Stop() 来触发优雅退出
}

// Addr 运行地址
func (conf *Config) Addr() string {
	return fmt.Sprintf("%s:%d", conf.IP, conf.Port)
}
