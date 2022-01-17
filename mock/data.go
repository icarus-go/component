package mock

import (
	"github.com/icarus-go/data/json"
	"net/url"
)

// JSON 模拟请求参数对象
type JSON struct {
	data map[string]interface{}
}

// NewJSON 新JSON参数对象
func NewJSON() *JSON {
	instance := new(JSON)

	instance.data = make(map[string]interface{})

	return instance
}

//NewJSONFormMap 根据Map转为JSON对象
func NewJSONFormMap(data map[string]interface{}) *JSON {
	instance := new(JSON)

	instance.data = data

	return instance
}

//NewJSONFormString
func NewJSONFormString(v string) (*JSON, error) {
	m := make(map[string]interface{})
	if err := json.Unmarshal([]byte(v), &m); err != nil {
		return nil, err
	}

	instance := new(JSON)
	instance.data = m
	return instance, nil
}

func NewJSONFormStruct(data interface{}) (*JSON, error) {
	marshalJSON, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	instance := new(JSON)

	instance.data = make(map[string]interface{})

	if err = json.Unmarshal(marshalJSON, &instance.data); err != nil {
		return nil, err
	}

	return instance, nil
}

// Add 添加参数
func (j *JSON) Add(key string, value interface{}) *JSON {
	j.data[key] = value
	return j
}

// Body 获取JSON字符串
func (j *JSON) Body() []byte {
	body, _ := json.Marshal(j.data)
	return body
}

type Parameter struct {
	values url.Values
}

func NewParameter() *Parameter {
	return &Parameter{
		values: make(map[string][]string),
	}
}

func (p *Parameter) Add(key string, value string) {
	p.values.Add(key, value)
}

func (p *Parameter) Encode() string {
	return p.values.Encode()
}

func (p *Parameter) Values() url.Values {
	return p.values
}
