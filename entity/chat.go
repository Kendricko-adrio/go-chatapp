package entity

import (
	"gorm.io/gorm"
)

type Chat struct {
	GroupID    uint
	Group      Group `gorm:"foreignKey:GroupID"`
	UserFromID uint
	UserFrom   User `gorm:"foreignKey:UserFromID"`
	Message    string
	gorm.Model
}
