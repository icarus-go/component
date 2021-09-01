package repo

import (
	"gorm.io/gorm"
	"time"
)

type S struct {
	ID        string         `json:"id" gorm:"id,primary" example:"ID"` // ID
	CreatedAt time.Time      `json:"createdAt,omitempty"`               // 创建时间
	UpdatedAt time.Time      `json:"updatedAt,omitempty"`               // 更新数据
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`                    // 删除时间
}

type M struct {
	ID        uint64         `json:"id" gorm:"id,primary" example:"1"` // ID
	CreatedAt time.Time      `json:"createdAt,omitempty"`              // 创建时间
	UpdatedAt time.Time      `json:"updatedAt,omitempty"`              // 更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`                   // 删除时间
}
