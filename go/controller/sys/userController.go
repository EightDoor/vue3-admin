package ControllerSys

import (
	"github.com/gin-gonic/gin"
	"zhoukai/model/sys"
	ServiceSys "zhoukai/service/sys"
	"zhoukai/utils"
)


// @Tags 用户
// @Description 获取用户列表
// @Param   page query	string false "page"
// @Param   size_size query	string false "page_size"
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
	singData, result := ServiceSys.UserSinge(id)
	utils.R(singData, result, c)
}
// @Tags 用户
// @Description 创建用户
// @Param account body string true "用户登录账号"
// @Param pass_word body string true "密码"
// @Param nick-name body string true "用户显示昵称"
// @Param avatar body string false "头像地址"
// @Param dept_id body string false "部门id"
// @Param email body string false "邮箱"
// @Param status body int false "所属状态是否有效  1是有效 0是失效"
// @Param phone_num body string false "用户手机号码"
// @Success 200 {string} string	"ok"
// @Router /api/v1/user [post]
func UserCreate(c *gin.Context)  {
	var sysUser ModelSys.SysUser
	data := utils.Verify(&sysUser, c)
	if data == nil {
		data, result := ServiceSys.UserCreate(sysUser)
		utils.R(data, result, c)
	}
}

// @Tags 用户
// @Description 修改用户
// @Param id path string true "id"
// @Param account body string true "用户登录账号"
// @Param pass_word body string true "密码"
// @Param nickName body string true "用户显示昵称"
// @Param avatar body string false "头像地址"
// @Param dept_id body string false "部门id"
// @Param email body string false "邮箱"
// @Param status body int false "所属状态是否有效  1是有效 0是失效"
// @Param phone_num body string false "用户手机号码"
// @Success 200 {string} string	"ok"
// @Router /api/v1/user/{id} [put]
func UserUpdate(c *gin.Context)  {
	var sysUser ModelSys.SysUser
	sysUser.ID = c.Param("id")
	valid := utils.Verify(&sysUser, c)
	if valid == nil {
		data, result := ServiceSys.UserUpdate(sysUser)
		utils.R(data, result, c)
	}
}
// @Tags 用户
// @Description 删除用户
// @Param id path string true "id"
// @Router /api/v1/user/{id} [delete]
func UserDel(c *gin.Context)  {
	var sysUser ModelSys.SysUser
	id := c.Param("id")
	sysUser.ID = id
	result := ServiceSys.UserDel(sysUser)
	utils.R(id, result, c)
}