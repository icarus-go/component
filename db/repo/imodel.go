package repo

import (
	"gorm.io/gorm"
	"time"
)

type S struct {
	ID        string         `json:"id" gorm:"column:id;size:32;primaryKey" swaggertype:"string" example:"uint64 主键ID"` // ID
	CreatedAt Datetime       `json:"createdAt,omitempty" swaggertype:"string" example:"创建时间"`                           // 创建时间
	UpdatedAt Datetime       `json:"updatedAt,omitempty" swaggertype:"string" example:"更新时间"`                           // 更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index" swaggerignore:"true"`                                               // 删除时间
}

func (s *S) BeforeCreate(tx *gorm.DB) error {
	now := time.Now().Local()
	s.CreatedAt.Time = now
	s.UpdatedAt.Time = now
	return nil
}

func (s *S) BeforeUpdate(tx *gorm.DB) error {
	now := time.Now().Local()
	s.UpdatedAt.Time = now
	return nil
}

type M struct {
	ID        uint64         `json:"id" gorm:"column:id;size:20;primaryKey" swaggertype:"string" example:"uint64 主键ID"` // ID
	CreatedAt Datetime       `json:"createdAt,omitempty" swaggertype:"string" example:"创建时间"`                           // 创建时间
	UpdatedAt Datetime       `json:"updatedAt,omitempty" swaggertype:"string" example:"更新时间"`                           // 更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index" swaggerignore:"true"`                                               // 删除时间
}

func (m *M) BeforeCreate(tx *gorm.DB) error {
	now := time.Now().Local()
	m.CreatedAt.Time = now
	m.UpdatedAt.Time = now
	return nil
}

func (m *M) BeforeUpdate(tx *gorm.DB) error {
	now := time.Now().Local()
	m.UpdatedAt.Time = now
	return nil
}
