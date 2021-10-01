package common

import "strings"

type Cos struct {
	Fields []string      `json:"fields" form:"fields"`
	Args   []interface{} `json:"args" form:"args"`
}

//Join 拼接
func (s *Cos) Join() string {
	return strings.Join(s.Fields, " , ")
}

func (s *Cos) Cond() interface{} {
	return s.Args
}
