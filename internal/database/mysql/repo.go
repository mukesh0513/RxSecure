package mysql
//
//import (
//	"github.com/gomodule/redigo/redis"
//	"github.com/jinzhu/gorm"
//	"github.com/mukesh0513/RxSecure/internal/database"
//	"github.com/mukesh0513/RxSecure/internal/model"
//	"log"
//)
//
//type MysqlConnection struct{}
//
//var mysqlConn *gorm.DB
//
//func Initialize(db *gorm.DB, logging bool) {
//	if mysqlConn == nil {
//		mysqlConn = db
//		db.LogMode(logging)
//		db.AutoMigrate(&model.Keys{})
//		database.DB = db
//	}
//}
//
//func (db *MysqlConnection) Find(table string, model interface{}, args interface{}) interface{} {
//	if err := mysqlConn.Where("token = ?", model.Token).Find(&model).Error; err != nil {
//		log.Println(err)
//	}
//	return model
//}
//
//func (db *MysqlConnection)Create(table string, model interface{}, args interface{}) interface{} {
//	return  mysqlConn.Create(&model)
//}
//
//func (db *MysqlConnection) Delete() {
//
//}
//
//func GetMysqlConn() *gorm.DB {
//	return mysqlConn
//}