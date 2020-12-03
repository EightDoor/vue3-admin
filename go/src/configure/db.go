package configure

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitMySqlCon() (err error) {
	// 可以在config包里设置成init函数
	ParserConfig()
	connStr  := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", DbConfig.User, DbConfig.Passwd, DbConfig.Host, DbConfig.Port, DbConfig.Database)
	fmt.Println(connStr)
	DB, err = gorm.Open(mysql.Open(connStr), &gorm.Config{})

	if err != nil {
		return err
	}
	return nil
}
