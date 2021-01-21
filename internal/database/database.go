package database

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/mukesh0513/RxSecure/internal/config"
	"github.com/mukesh0513/RxSecure/internal/model"
	redisDb "github.com/mukesh0513/RxSecure/internal/provider/redis"
)

var (
	DB  *gorm.DB
	err error
)

type Database struct {
	*gorm.DB
}

// Setup opens a database and saves the reference to `Database` struct.
func Setup() {
	config := config.GetConfig()

	driver := config.Database.Driver
	database := config.Database.Dbname
	username := config.Database.Username
	password := config.Database.Password
	host := config.Database.Host
	port := config.Database.Port

	switch driver {
	case "mysql":
		db, err := gorm.Open("mysql", username+":"+password+"@tcp("+host+":"+port+")/"+database+"?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			fmt.Println("db err: ", err)
			break
		}

		db.LogMode(config.Database.LogMode)
		db.AutoMigrate(&model.Keys{})
		db.AutoMigrate(&model.Data{})
		DB = db

	case "redis":
		conn, err := redis.Dial("tcp", host+":"+port)
		if err != nil {
			fmt.Println("db err: ", err)
			break
		}
		redisDb.Initialize(conn)
	}
}

// GetDB helps you to get a connection
func GetDB() *gorm.DB {
	return DB
}

//func SelectDbAndFire()  {
//	mysqlConnection := mysql.GetMysqlConn()
//	redisConnection := redisDb.GetRedisConn()
//	switch true {
//	case mysqlConnection == nil:
//		mysql.MysqlConnection{}.Find()
//	case redisConnection == nil:
//
//	}
//}