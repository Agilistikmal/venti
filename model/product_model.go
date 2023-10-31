package model

import "gorm.io/gorm"

type Product struct {
	*gorm.Model
	ID          string `gorm:"primarykey,unique"`
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	Price       int    `gorm:"not null"`
	Stars       int    `gorm:"not null,default=0"`
	Sold        int    `gorm:"not null,default=0"`
	Stock       int    `gorm:"not null,default=0"`
	Usage       string
}
