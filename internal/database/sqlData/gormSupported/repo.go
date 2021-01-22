package gormSupported

import (
	"github.com/jinzhu/gorm"
	"github.com/mukesh0513/RxSecure/internal/model"
)

type GormConnection struct{}

var (
	gormConn         *gorm.DB
	//GormConnProvider sqlData.ISqlDatabase
)

func Initialize(db *gorm.DB, logging bool) {
	if gormConn == nil {
		gormConn = db
		//GormConnProvider = GormConnection{}
		db.LogMode(logging)
		db.AutoMigrate(&model.EncKeys{})
		db.AutoMigrate(&model.Data{})
	}
}

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

//func (db GormConnection) Find(model interface{}, args interface{}) (interface{}, error) {
//
//}
//
//func (db GormConnection) Create(model interface{}) error {
//
//}
//
//func (db GormConnection) Delete(model interface{}, args interface{}) error {
//	var err error
//	err = gormConn.Where(args).Delete(&model).Error
//
//	return err
//}

//func GetMysqlConn() *gorm.DB {
//	return gormConn
//}
