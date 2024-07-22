package entities

import (
	"gorm.io/gorm"
	"time"
)

type Artist struct {
	ID        uint           `gorm:"primary_key"`
	Name      string         `gorm:"type:varchar(100);not null"`
	Bio       string         `gorm:"type:text"`
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
