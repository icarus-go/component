package arg

import (
	"flag"
	"pmo-test4.yz-intelligence.com/kit/component/gins"
)

var (
	Name    string
	Host    string
	Ip      string
	Port    int
	Timeout int
	Debug   string
	Pprof   bool
	Cors    bool
	Config  *gins.Config
)

func init() {
	flag.StringVar(&Name, "name", "", "应用名称")
	flag.StringVar(&Host, "host", "", "host名称")
	flag.StringVar(&Ip, "ip", "", "运行ip 默认0.0.0.0")
	flag.IntVar(&Port, "port", 8080, "")
	flag.IntVar(&Timeout, "timeout", 30, "退出时间")
	flag.StringVar(&Debug, "debug", "debug", "调试模式")
	flag.BoolVar(&Pprof, "pprof", false, "性能监控模式")
	flag.BoolVar(&Cors, "cors", false, "是否")
	flag.Parse()
	if Ip == "" {
		Ip = "0.0.0.0"
	}
	if Port <= 0 {
		Port = 8080
	}
	if Timeout <= 0 {
		Timeout = 30
	}
	Name = "yz.demo"
}
