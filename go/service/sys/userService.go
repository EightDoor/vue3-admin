package ServiceSys

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"zhoukai/configure"
	"zhoukai/model/sys"
	"zhoukai/utils/db"
)

func UserQuery(dates []ModelSys.SysUser, c *gin.Context)([]ModelSys.SysUser, *gorm.DB)  {
	result := configure.DB.Scopes(UtilsDB.Paginate(c)).Find(&dates)
	return dates, result
}
func UserSinge(id string)(ModelSys.SysUser, *gorm.DB)  {
	var data ModelSys.SysUser
	result := configure.DB.First(&data, id)
	return data, result
}
func UserCreate(data ModelSys.SysUser)(ModelSys.SysUser, *gorm.DB)  {
	result := configure.DB.Scopes(UtilsDB.CreateUUId()).Create(&data)
	fmt.Println(result)
	return data, result
}


