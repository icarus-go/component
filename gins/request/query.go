package request

import "errors"

type Query struct {
	String struct {
		ID  string   `json:"id" example:"ID" v:"required|length:1,1000#请输入id|id长度为:min到:max位"` // ID,当前业务主键
		IDs []string `json:"ids" example:"ID组"`                                                // ID串组
	} `json:"string"`

	Integer struct {
		ID  uint64   `json:"id" example:"1"` // ID
		IDs []uint64 `json:"ids"`            // ID串组
	} `json:"integer"` // 整形ID

	Page     int `json:"page" form:"page" example:"1"`          // 页码
	PageSize int `json:"pageSize" form:"pageSize" example:"20"` // 页面最大条数
}

//ResetPage 重置
func (q *Query) ResetPage() *Query {
	if q.Page < 1 {
		q.Page = 1
	}
	if q.PageSize < 20 {
		q.PageSize = 20
	}
	return q
}

func (q *Query) HasIntegerID() error {
	if q.Integer.ID == 0 {
		return errors.New("ID不允许为0")
	}
	return nil
}

func (q *Query) HasIntegerIDs() error {
	if len(q.Integer.IDs) < 1 {
		return errors.New("IDs不允许为空")
	}
	return nil
}

func (q *Query) HasStringIDValue() error {
	if q.String.ID == "" {
		return errors.New("ID不允许为空")
	}
	return nil
}

func (q *Query) HasStringIDsValue() error {
	if len(q.String.IDs) < 1 {
		return errors.New("ids长度必须大于0")
	}
	return nil
}
