package ControllerSys

import (
	"github.com/gin-gonic/gin"
	ModelSys "zhoukai/model/sys"
	ServiceSys "zhoukai/service/sys"
	"zhoukai/utils"
)
var sysDict ModelSys.SysDict

// Tags 字典
// @Description 字典列表
// @Param  name body string true "字典名称"
// @Param  serial_number body string false "字典编号"
// @Param  describe body string false "描述"
// @Router /api/v1/dict [post]
func DictCreate(c *gin.Context)  {
	valida := utils.Verify(&sysDict, c)
	if valida == nil {
		data, result := ServiceSys.DictCreate(sysDict)
		utils.R(data, result, c)
	}
}
// Tags 字典
// @Description 字典列表
// @Param page query string false "page"
// @Param page_size query string false "page_size"
// @Router /api/v1/dict [get]
func DictList(c *gin.Context)  {
	var sysDictList []ModelSys.SysDict
	data, result := ServiceSys.DictList(sysDictList,c)
	utils.R(data, result, c)
}
// @Tags 字典
// @Description 字典修改
// @Param  name body string true "字典名称"
// @Param  serial_number body string false "字典编号"
// @Param  describe body string false "描述"
// @Param  id  path string true "主键"
// @Router /api/v1/dict/{id} [put]
func DictEdit(c * gin.Context)  {
	sysDict.ID = c.Param("id")
	validate := utils.Verify(&sysDict, c)
	if validate == nil {
		data, result := ServiceSys.DictUpdate(sysDict)
		utils.R(data, result, c)
	}
}
// @Tags 字典
// @Description 字典删除
// @Param id path string true "主键"
// Router /api/v1/dict/{id} [delete]
func DictDel(c *gin.Context)  {
	sysDict.ID = c.Param("id")
	data, result := ServiceSys.DictDel(sysDict)
	utils.R(data, result, c)
}