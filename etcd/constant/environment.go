package constant

type Environment string

//Equals 是否相等
func (e Environment) Equals(environment string) bool {
	return e.Value() == environment
}

//Value 值
func (e Environment) Value() string {
	return string(e)
}
