package common

import (
	"errors"
	"strings"
)

type (
	Query struct {
		String  String  `json:"string" form:"string"`   // 字符串ID
		Integer Integer `json:"integer" form:"integer"` // 整形ID
		Paging
	}
	String struct {
		ID  string   `json:"id" form:"id" example:"ID" v:"required|length:1,1000#请输入id|id长度为:min到:max位"` // ID,当前业务主键
		IDs []string `json:"ids" form:"ids" example:"ID组"`                                               // ID串组
	}

	Integer struct {
		ID  uint64   `json:"id" example:"1" form:"id"` // ID
		IDs []uint64 `json:"ids" form:"ids"`           // ID串组
	}

	Paging struct {
		Page     int `json:"page" form:"page" example:"1"`          // 页码
		PageSize int `json:"pageSize" form:"pageSize" example:"20"` // 页面最大条数
	}

	Order struct {
		Fields []string `json:"fields" form:"fields" example:"id,createTime"` // 排序字段
		Type   string   `json:"type" form:"type" example:"desc/asc"`          // 排序类型
	}

	GetByID struct {
		ID uint64 `json:"id" form:"id" swggertype:"string" example:"uint64 主键ID"` // 主键ID
	}

	GetByIDs struct {
		IDs []uint64 `json:"ids" form:"ids" swggertype:"string" example:"[]uint64 主键IDs"` // 主键IDs
	}
)

//ResetPage 重置
func (q *Query) ResetPage() *Query {
	if q.Page < 1 {
		q.Paging.Page = 1
	}
	if q.PageSize < 20 {
		q.PageSize = 20
	}
	return q
}

func (q *Integer) HasID() error {
	if q.ID == 0 {
		return errors.New("ID不允许为0")
	}
	return nil
}

func (q *Integer) HasIDs() error {
	if len(q.IDs) < 1 {
		return errors.New("IDs不允许为空")
	}
	return nil
}

func (q *String) HasID() error {
	if q.ID == "" {
		return errors.New("ID不允许为空")
	}
	return nil
}

func (q *String) HasIDs() error {
	if len(q.IDs) < 1 {
		return errors.New("ids长度必须大于0")
	}
	return nil
}

func (o *Order) Join() string {
	return strings.Join(o.Fields, ",")
}

func (o *Order) Sort() string {
	if o.Type == "" {
		return ""
	} // 如果为空，默认使用降序
	if o.Type == "desc" || o.Type == "asc" {
		return o.Type
	} // 符合数据库值，返回
	return "" // 不符合，直接返回空，按照升序处理

}
