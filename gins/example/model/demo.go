package model

// Demo 示例
type Demo struct {
	ID        string `json:"id"`   // ID
	Name      string `json:"name"` // 名字
	FirstName string `json:"firstName" xorm:"'first_name'"`
	CreateAt  int64  `json:"createAt"` // 创建时间
	UpdateAt  int64  `json:"updateAt"` // 最后更新时间
}
