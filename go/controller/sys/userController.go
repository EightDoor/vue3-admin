package ControllerSys

import (
	"github.com/gin-gonic/gin"
	"zhoukai/model/sys"
	"zhoukai/service"
	ServiceSys "zhoukai/service/sys"
	"zhoukai/utils"
)

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
// @Tags 用户
// @Description 获取单条用户信息
// @Param   id path	string true "id"
// @Success 200 {string} string	"ok"
// @Router /api/v1/user/{id} [get]
func UserSinge(c *gin.Context)  {
	id := c.Param("id")
	data := ModelSys.SysUser{
		Id: id,
	}
	singData, result := ServiceSys.UserSinge(data)
	utils.R(singData, result, c)
}
// @Tags 用户
// @Description 创建用户
// @Success 200 {string} string "ok"
// @Param Account body string true "用户登录账号"
// @Param PassWord body string true "密码"
// @Param NickName body string true "用户显示昵称"
// @Param Avatar body string false "头像地址"
// @Param DeptId body string false "部门id"
// @Param Email body string false "邮箱"
// @Param Status body int false "所属状态是否有效  1是有效 0是失效"
// @Param PhoneNum body string false "用户手机号码"
// @Router /api/v1/user [post]
func UserCreate(c *gin.Context)  {
	var sysUser ModelSys.SysUser
	data := utils.Verify(&sysUser, c)
	if data == nil {
		data, result := ServiceSys.UserCreate(sysUser)
		utils.R(data, result, c)
	}
}