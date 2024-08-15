package model

import "gorm.io/gorm"

type Demo struct {
	gorm.Model
	Demoname string `gorm:"not null"`
}

func (u *Demo) TableName() string {
	return "demos"
}
