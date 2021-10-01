package common

import "strings"

type Cos struct {
	Fields []string      `json:"fields"`
	Args   []interface{} `json:"args"`
}

//Join 拼接
func (s *Cos) Join() string {
	return strings.Join(s.Fields, " , ")
}

func (s *Cos) Cond() interface{} {
	return s.Args
}
