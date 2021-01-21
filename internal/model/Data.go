package model

import "github.com/jinzhu/gorm"

type Data struct {
	gorm.Model
	Token   string `redis:"token" json:"token"  gorm:"column:token"`
	Payload string `redis:"payload" json:"payload"  gorm:"column:payload"`
}
