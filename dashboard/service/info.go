package service

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

// ServiceName 服务名称
func (info *Info) ServiceName(name string) *Info {
	info.bill["name"] = name
	return info
}

func (info *Info) Version(version string) *Info {
	info.bill["version"] = version
	return info
}
