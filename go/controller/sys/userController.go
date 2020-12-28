package ControllerSys

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zhoukai/configure"
	"zhoukai/middleware"
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

// @Tags 用户
// @Descriptions 查询拥有的菜单
// @Param id path string "用户id"
// @Router /api/v1/user/permissions/{id} [get]
func UserPermissions(c *gin.Context)  {
	var datas []ModelSys.SysRole
	id := c.Param("id")
	data, result := ServiceSys.UserPermission(id, datas,c)
	utils.R(data, result, c)
}
// @Tags 用户
// @Descriptions 用户拥有菜单
// @Router /api/v1/user/permissions	[post]
func UserAssociatedMenu(c *gin.Context)  {
	var userRole ModelSys.SysUserRole
	err := utils.Verify(&userRole, c)
	if err == nil {
		data, result := ServiceSys.UserAssociatedMenu(userRole)
		utils.R(data, result, c)
	}
}

// Tags 用户
// @Descriptions 获取用户信息(登录成功之后header携带 token参数获取)
// @Router /api/v1/user/getInfo [get]
func GetUserInfo(c *gin.Context)  {
	j := middleware.NewJWT()
	token := c.Request.Header.Get("token")
	if token == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": configure.RequestError,
			"msg":    "请求未携带token，无权限访问",
			"data":   nil,
		})
		c.Abort()
		return
	}
	result, error := j.ParserToken(token)
	if error ==nil {
		var sysUserInfo ModelSys.SysUser
		sysUserInfo.ID = result.ID
		data, r := ServiceSys.UserGetInfo(sysUserInfo)
		utils.R(data, r ,c)
	}
}

// @Tags 用户
// @Descriptions 获取用户拥有的菜单、角色
// @Router /api/v1/user/getWithTheMenu [get]
func UserGetWithMenu(c *gin.Context)  {
	j := middleware.NewJWT()
	token := c.Request.Header.Get("token")
	result, error := j.ParserToken(token)
	if error ==nil {
		var userWithTheMenu ModelSys.SysUserWithTheMenu
		var id string
		id = result.ID
		data, r := ServiceSys.UserGetWithMenu(id, userWithTheMenu)
		utils.R(data, r ,c)
	}
}