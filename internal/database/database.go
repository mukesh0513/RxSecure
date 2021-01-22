package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/mukesh0513/RxSecure/internal/config"
	"github.com/mukesh0513/RxSecure/internal/database/sqlData/gormSupported"
)

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
		gormSupported.Initialize(db, config.Database.LogMode)

	case "redis":
		//conn, err := redis.Dial("tcp", host+":"+port)
		//if err != nil {
		//	fmt.Println("db err: ", err)
		//	break
		//}
		//redis2.Initialize(conn)
	}
}