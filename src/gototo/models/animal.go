package models

import "github.com/jinzhu/gorm"

type Animal struct {

	gorm.Model

	Name string `gorm:"size:255"`
	Group string `gorm:"size:255"`
}
