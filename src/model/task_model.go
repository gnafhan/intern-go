package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	ID          uuid.UUID `gorm:"primaryKey;not null" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Description string    `gorm:"not null" json:"description"`
	CategoryID  uuid.UUID `gorm:"not null" json:"category_id"`
	Category    Category  `gorm:"foreignKey:CategoryID" json:"category"`
	Priority    string    `gorm:"not null" json:"priority" enums:"low,medium,high"`
	Deadline    time.Time `gorm:"not null" json:"deadline"`
	CreatedAt   time.Time `gorm:"autoCreateTime:milli" json:"-"`
	UpdatedAt   time.Time `gorm:"autoCreateTime:milli;autoUpdateTime:milli" json:"-"`
}

func (task *Task) BeforeCreate(_ *gorm.DB) error {
	if task.ID == uuid.Nil {
		task.ID = uuid.New() // Generate UUID before create
	}
	return nil
}
