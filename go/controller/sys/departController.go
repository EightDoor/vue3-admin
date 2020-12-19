package ControllerSys

import (
	"github.com/gin-gonic/gin"
	ModelSys "zhoukai/model/sys"
	ServiceSys "zhoukai/service/sys"
	"zhoukai/utils"
)

// @Tags 部门
// @Descriptions 部门列表
// @Param page query string false "page"
// @Param page_size query string false "page_size"
// @Success 200 {string} string	"ok"
// @Router	/api/v1/depart [get]
func DepartList(c *gin.Context)  {
	var list []ModelSys.SysDept
	dates, result := ServiceSys.DepartList(list, c)
	utils.R(dates, result, c)
}

// @Tags 部门
// @Descriptions 添加部门
// @Param parent_id body string true "父级id"
// @Param name body string true "节点名称"
// @Param order_num body string true "排序"
// @Router /api/v1/depart [post]
func DepartCreate(c *gin.Context)  {
	var sysDepart ModelSys.SysDept
	valida := utils.Verify(&sysDepart, c)
	if valida == nil {
		data, result := ServiceSys.DepartCreate(sysDepart)
		utils.R(data, result, c)
	}
}
// @Tags 部门
// @Description 修改部门
// Param id path string true "id"
// @Param parent_id body string true "父级id"
// @Param name body string true "节点名称"
// @Param order_num body string true "排序"
// @Router /api/v1/depart/{id}	[put]
func DepartUpdate(c *gin.Context)  {
	var sysDepart ModelSys.SysDept
	valid := utils.Verify(&sysDepart, c)
	sysDepart.ID = c.Param("id")
	if valid == nil {
		data, result := ServiceSys.DepartUpdate(sysDepart)
		utils.R(data, result, c)
	}
}

// @Tags 部门
// @Description 删除部门
// Param id path string true "id"
// @Router /api/v1/depart/{id} [delete]
func DepartDel(c *gin.Context)  {
	var sysDepart ModelSys.SysDept
	sysDepart.ID = c.Param("id")
	data, result := ServiceSys.DepartDel(sysDepart)
	utils.R(data, result, c)
}
