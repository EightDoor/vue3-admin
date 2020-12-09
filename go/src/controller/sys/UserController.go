package ControllerSys

import (
	"github.com/gin-gonic/gin"
	"zhoukai/src/model"
	ServiceSys "zhoukai/src/service/sys"
	"zhoukai/src/utils"
)

func UserList(c *gin.Context)  {
	var list []model.SysUser
	dates, result := ServiceSys.UserQuery(list, c)
	utils.R(dates, result, c)
}