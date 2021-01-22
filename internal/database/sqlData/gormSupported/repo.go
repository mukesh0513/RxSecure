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
