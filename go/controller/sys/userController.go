package ControllerSys

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"zhoukai/model/sys"
	ServiceSys "zhoukai/service/sys"
	"zhoukai/utils"
)

// 用户 godoc
// @Summary Add a new pet to the store
// @Tags 用户
// @Description 获取用户列表
// @Param   page query	string false "page"
// @Param   size query	string false "size"
// @Success 200 {string} string	"ok"
// @Router /api/v1/user [get]
func UserList(c *gin.Context)  {
	var list []ModelSys.SysUser
	dates, result := ServiceSys.UserQuery(list, c)
	utils.R(dates, result, c)
}
func UserSinge(c *gin.Context)  {
	id, _ := strconv.Atoi(c.Param("id"))
	data := ModelSys.SysUser{
		Id: id,
	}
	singData, result := ServiceSys.UserSinge(data)
	utils.R(singData, result, c)
}
func UserCreate(c *gin.Context)  {
	var sysUser ModelSys.SysUser
	err := utils.Verify(&sysUser, c)
	if err {
		data, result := ServiceSys.UserCreate(sysUser)
		utils.R(data, result, c)
	}
}