package configure

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	DB *gorm.DB
)

func InitMySqlCon() (err error) {
	// 可以在config包里设置成init函数
	ParserConfig()
	connStr  := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", DbConfig.User, DbConfig.Passwd, DbConfig.Host, DbConfig.Port, DbConfig.Database)
	fmt.Println(connStr)
	DB, err = gorm.Open(mysql.Open(connStr), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	})

	if err != nil {
		return err
	}
	return nil
}
