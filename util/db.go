package util

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DB = InitDB()

func InitDB() *gorm.DB {
	//db_dsn位于util.static中
	db, err := gorm.Open(mysql.Open(DB_dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //全局日志：info
	})

	if err != nil {
		log.Println("gorm init error" + err.Error())
	}

	return db
}
