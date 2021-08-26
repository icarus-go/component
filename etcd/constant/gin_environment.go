package constant

const (
	//Develop 测试环境变量
	Develop Environment = "ETCD_DEVELOP_ADDR"
	//Release 正式环境变量
	Release Environment = "ETCD_RELEASE_ADDR"
	//Test 测试环境变量
	Test Environment = "ETCD_TEST_ADDR"
)

type Split string

const (
	// Comma 默认的切割符
	Comma Split = ","
)

//Value 值
func (c Split) Value() string {
	return string(c)
}
