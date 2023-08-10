package util

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"hasdsd.cn/op-dl-server/dao"
	"log"
)

var DB = dao.Query{}

func InitDB() {
	//db_dsn位于util.static中
	db, err := gorm.Open(mysql.Open(DB_dsn), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),        //全局日志：info
		DisableForeignKeyConstraintWhenMigrating: true,                                       //自动建表不会创建外键
		NamingStrategy:                           schema.NamingStrategy{SingularTable: true}, //自动建表配置表名
	})
	if err != nil {
		log.Println("gorm init error" + err.Error())
	}

	//关联表
	dao.SetDefault(db)
	q := dao.Q
	DB = *q
}
func init() {
	InitDB()
}
