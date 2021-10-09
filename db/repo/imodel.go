package repo

import "gorm.io/gorm"

type S struct {
	ID        string         `json:"id" gorm:"column:id;size:32;primaryKey" swaggertype:"string" example:"uint64 主键ID"` // ID
	CreatedAt Datetime       `json:"createdAt,omitempty" swaggertype:"string" example:"创建时间"`                           // 创建时间
	UpdatedAt Datetime       `json:"updatedAt,omitempty" swaggertype:"string" example:"更新时间"`                           // 更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index" swaggerignore:"true"`                                               // 删除时间
}

type M struct {
	ID        uint64         `json:"id" gorm:"column:id;size:20;primaryKey" swaggertype:"string" example:"uint64 主键ID"` // ID
	CreatedAt Datetime       `json:"createdAt,omitempty" swaggertype:"string" example:"创建时间"`                           // 创建时间
	UpdatedAt Datetime       `json:"updatedAt,omitempty" swaggertype:"string" example:"更新时间"`                           // 更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index" swaggerignore:"true"`                                               // 删除时间
}
