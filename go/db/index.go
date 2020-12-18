package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
	"zhoukai/configure"
)

var (
	DB *gorm.DB
)


func InitMySqlCon() (err error) {
	// logger
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      false,         // 禁用彩色打印
		})

	// 可以在config包里设置成init函数
	configure.ParserConfig()
	connStr  := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", configure.DbConfig.User, configure.DbConfig.Passwd, configure.DbConfig.Host, configure.DbConfig.Port, configure.DbConfig.Database)
	fmt.Println(connStr)
	DB, err = gorm.Open(mysql.Open(connStr), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
		Logger: newLogger,
	})

	if err != nil {
		return err
	}
	return nil
}
