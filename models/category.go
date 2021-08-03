package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Name        string `json:"name" gorm:"type:text"`
	Description string `json:"description" gorm:"type:text"`
}
