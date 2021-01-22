package gormSupported

import (
	"github.com/mukesh0513/RxSecure/internal/model"
	"log"
)

type PayloadService struct{}

func (service PayloadService) GetEncryptedData(token string) string {
	var data model.Data

	if err := gormConn.Where("token = ?", token).Find(&data).Error; err != nil {
		log.Println(err)
	}
	return data.Payload
}

func (service PayloadService) SetEncryptedData(token string, payload string) string {
	var model model.Data
	model.Token= token
	model.Payload = payload
	gormConn.Create(&model)

	return ""
}

func NewPayloadService() PayloadService {
	return PayloadService{}
}
