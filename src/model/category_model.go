package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID        uuid.UUID `gorm:"primaryKey;not null" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	CreatedAt time.Time `gorm:"autoCreateTime:milli" json:"-"`
	UpdatedAt time.Time `gorm:"autoCreateTime:milli;autoUpdateTime:milli" json:"-"`
}

func (category *Category) BeforeCreate(_ *gorm.DB) error {
	if category.ID == uuid.Nil {
		category.ID = uuid.New() // Generate UUID before create
	}
	return nil
}
