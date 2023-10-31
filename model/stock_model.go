package model

import "gorm.io/gorm"

type Stock struct {
	*gorm.Model
	ProductID string `gorm:"not null"`
	Content   string `gorm:"not null"`
}
