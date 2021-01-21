package model

import "github.com/jinzhu/gorm"

type Data struct {
	gorm.Model
	Token  string `json:"token"  gorm:"column:token"`
	Payload string `json:"payload"  gorm:"column:payload"`
}
