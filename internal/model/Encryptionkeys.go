package model

import "github.com/jinzhu/gorm"

type EncKeys struct {
	gorm.Model
	EncKey string `json:"enc_key"  gorm:"column:enc_key"`
}
