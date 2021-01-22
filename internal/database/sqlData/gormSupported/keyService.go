package gormSupported

import "github.com/mukesh0513/RxSecure/internal/model"

type KeyService struct{}

func (service KeyService) GetEncryptedKey(hash int64) string {
	var model model.EncKeys
	gormConn.Where("id = ?", hash).Find(&model)

	return model.EncKey
}

func (service KeyService) SetEncryptedKey(hash int64, enc_key string) string {
	var model model.EncKeys
	model.EncKey = enc_key
	gormConn.Create(&model)

	return ""
}

func NewKeyService() KeyService {
	return KeyService{}
}
