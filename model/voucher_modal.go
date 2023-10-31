package model

import (
	"gorm.io/gorm"
	"time"
)

type Voucher struct {
	*gorm.Model
	Code      string `gorm:"not null,unique"`
	Discount  int    `gorm:"not null"`
	ExpiredAt time.Time
}
