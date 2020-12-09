package ServiceSys

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"zhoukai/src/configure"
	"zhoukai/src/model"
	response_data "zhoukai/src/response-data"
)

// 查询用户列表
func UserQuery(dates []model.SysUser, c *gin.Context)([]model.SysUser, *gorm.DB)  {
	offset, pageSize, _ := response_data.PageIngService(c)
	result := configure.DB.Order("id desc").Offset(offset).Limit(pageSize).Find(&dates)
	return dates, result
}

