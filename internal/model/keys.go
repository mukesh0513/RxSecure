package model

import "github.com/jinzhu/gorm"

type Keys struct {
	gorm.Model
	Token 	  string `json:"token"  gorm:"column:token"`
}
