package ControllerSys

import (
	"github.com/gin-gonic/gin"
	ModelSys "zhoukai/model/sys"
	ServiceSys "zhoukai/service/sys"
	"zhoukai/utils"
)
var sysMenu ModelSys.SysMenu
// @Tags 菜单
// @Description 列表
// @Param page query int false "page"
// @Param page_size query int false "page_size"
// @Router /api/v1/menu [get]
func MenuList(c *gin.Context)  {
	var sysMenus []ModelSys.SysMenu
	datas, result := ServiceSys.MenuList(sysMenus, c)
	utils.R(datas, result, c)
}

// @Tags 菜单
// @Description 创建
// @Param parent_id body string true "父级id"
// @Param name body string true "菜单名称"
// @Param type body body int true "类型 1. 菜单/目录 2 tabs 3 按钮"
// @Param order_num body int true "排序"
// @Param perms body string false "权限标识，接口标识"
// @Param code body string true "菜单标识，前端路由name"
// @Router /api/v1/menu [post]
func MenuCreate(c *gin.Context)  {
	err := utils.Verify(&sysMenu, c)
	if err == nil {
		data, result := ServiceSys.MenuCreate(sysMenu)
		utils.R(data, result, c)
	}
}

// @Tags 菜单
// @Description 编辑
// @Param id path string true "主键"
// @Param parent_id body string true "父级id"
// @Param name body string true "菜单名称"
// @Param type body body int true "类型 1. 菜单/目录 2 tabs 3 按钮"
// @Param order_num body int true "排序"
// @Param perms body string false "权限标识，接口标识"
// @Param code body string true "菜单标识，前端路由name"
// @Router /api/v1/menu/{id} [put]
func MenuUpdate(c *gin.Context)  {
	sysMenu.ID = c.Param("id")
	err := utils.Verify(&sysMenu, c)
	if err == nil {
		data, result := ServiceSys.MenuUpdate(sysMenu)
		utils.R(data, result, c)
	}
}
