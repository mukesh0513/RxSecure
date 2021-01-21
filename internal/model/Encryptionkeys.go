package model

import "github.com/jinzhu/gorm"

type Keys struct {
	gorm.Model
	EncKey string `redis:"enc_key" json:"enc_key"  gorm:"column:enc_key"`
}
