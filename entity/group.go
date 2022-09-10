package entity

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	GroupName string
	GroupType string
}

type GroupDetail struct {
	UserID  uint
	GroupID uint
	Group   Group `gorm:"foreignKey:GroupID"`
	User    User  `gorm:"foreignKey:UserID"`
}
