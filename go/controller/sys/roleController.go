package ControllerSys

import (
	"github.com/gin-gonic/gin"
	ModelSys "zhoukai/model/sys"
	ServiceSys "zhoukai/service/sys"
	"zhoukai/utils"
)
var sysRole ModelSys.SysRole

// @Tags 角色
// @Descriptions 角色列表
// @Param page query string false "page"
// @Param page_size query string false "page_size"
// @Router /api/v1/role [get]
func RoleList(c *gin.Context)  {
	var sysRoleList []ModelSys.SysRole
	datas, result := ServiceSys.RoleList(sysRoleList, c)
	utils.R(datas, result, c)
}

// @Tags 角色
// @Descriptions 创建角色
// @Param remark body string false "角色备注"
// @Param role_name body string true "角色名称"
// @Router /api/v1/role [post]

func RoleCreate(c *gin.Context)  {
	err := utils.Verify(&sysRole, c)
	if err == nil {
		data, result := ServiceSys.RoleCreate(sysRole)
		utils.R(data, result, c)
	}
}

// @Tags 角色
// @Descriptions 修改
// @Param id path string true "主键"
// @Param remark body string false "角色备注"
// @Param role_name body string true "角色名称"
// @Router /api/v1/role/{id} [put]
func RoleUpdate(c *gin.Context)  {
	err := utils.Verify(&sysRole, c)
	sysRole.ID = c.Param("id")
	if err == nil {
		data, result := ServiceSys.RoleUpdate(sysRole)
		utils.R(data, result, c)
	}
}

// @Tags 角色
// @Descriptions 删除
// @Param id path string "主键"
// @Router /api/v1/role/{id} [delete]
func RoleDel(c *gin.Context)  {
	sysRole.ID = c.Param("id")
	data, result := ServiceSys.RoleDel(sysRole)
	utils.R(data, result, c)
}