package errors

import "errors"

var (
	//Empty 空
	Empty = errors.New("environment is empty")
	//Invalid 无效值
	Invalid = errors.New("environment value is invalid")
	//ValueEmpty 环境变量值为空
	ValueEmpty = errors.New("environment value is empty")
)

type EnvironmentError struct {
	VariableValue string
	VariableName  string
	Err           error
	AppendErr     error
}

func (e *EnvironmentError) Error() string {
	if e == nil {
		return "<nil>"
	}

	result := ""
	if e.Err == Empty {
		result += "check your environment"
	} else if e.Err == Invalid {
		result += "environment variable name :" + e.VariableName + " ，value :" + e.VariableValue
	} else if e.Err == ValueEmpty {
		result += "environment variable name : " + e.VariableName + " , but the value is empty"
	} else {
		result += "environment variable name :" + e.VariableName + " ，value :" + e.VariableValue
	}
	result += e.Err.Error()
	return result
}

type ConfigError struct {
	Step string
	Err  error
}
