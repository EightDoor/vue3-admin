package ServiceSys

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"zhoukai/configure"
	"zhoukai/model/sys"
	"zhoukai/utils"
)

func UserQuery(dates []ModelSys.SysUser, c *gin.Context)([]ModelSys.SysUser, *gorm.DB)  {
	offset, pageSize, _ := utils.PageIngService(c)
	result := configure.DB.Offset(offset).Limit(pageSize).Find(&dates)
	return dates, result
}
func UserSinge(data ModelSys.SysUser)(ModelSys.SysUser, *gorm.DB)  {
	result := configure.DB.First(&data, data.Id)
	return data, result
}
func UserCreate(data ModelSys.SysUser)(ModelSys.SysUser, *gorm.DB)  {
	result := configure.DB.Create(&data)
	return data, result
}


