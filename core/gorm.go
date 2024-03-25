package core

import (
	"fmt"
	"go-blog-server/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

func InitGorm() {
	if global.CONFIG.Mysql.Host == "" {
		fmt.Println("未配置mysql，退出链接 GORM")
	}

	dsn := global.CONFIG.Mysql.Dsn()
	var mysqlLogger logger.Interface
	if global.CONFIG.System.Env == "debug" {
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Error)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		log.Fatalln(err)
	}

	sqlDB, _ := db.DB()
	//最大空闲连接处
	sqlDB.SetMaxIdleConns(10)
	//最大可容纳连接数
	sqlDB.SetMaxOpenConns(100)
	//最大可复用时间，不能超过 mysql wait_time
	sqlDB.SetConnMaxLifetime(time.Hour * 4)

	global.DB = db
}
