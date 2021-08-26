package constant

const (
	//DevelopF 测试环境变量
	DevelopF Environment = "develop"
	//ReleaseF 正式环境变量
	ReleaseF Environment = "release"
	//TestF 测试环境变量
	TestF Environment = "test"
)

type fileSystemMode string

const (
	//FileEnvironment 文件类型如何获取环境配置 系统变量名
	FileEnvironment fileSystemMode = "FILE_ENVIRONMENT"
)

//Value 值对象
func (f fileSystemMode) Value() string { return string(f) }
