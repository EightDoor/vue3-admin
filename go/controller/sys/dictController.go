package ControllerSys

import (
	"github.com/gin-gonic/gin"
	ModelSys "zhoukai/model/sys"
	ServiceSys "zhoukai/service/sys"
	"zhoukai/utils"
)
var sysDict ModelSys.SysDict
var sysDictItem ModelSys.SysDictItem

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


// Tags 字典
// @Description 字典项列表
// @Param  value body string true "数据值"
// @Param  label body string false "名称"
// @Param  describe body string false "描述"
// @Param  dict_id body string true "关联key"
// @Router /api/v1/dict-item [post]
func DictItemCreate(c *gin.Context)  {
	valida := utils.Verify(&sysDictItem, c)
	if valida == nil {
		data, result := ServiceSys.DictItemCreate(sysDictItem)
		utils.R(data, result, c)
	}
}
// Tags 字典
// @Description 字典项列表
// @Param page query string false "page"
// @Param page_size query string false "page_size"
// @Router /api/v1/dict-item [get]
func DictItemList(c *gin.Context)  {
	var sysDictItemList []ModelSys.SysDictItem
	id := c.Param("id")
	data, result := ServiceSys.DictItemList(sysDictItemList,c, id)
	utils.R(data, result, c)
}
// @Tags 字典
// @Description 字典项修改
// @Param  value body string true "字典值"
// @Param  label body string false "字典项"
// @Param  describe body string false "描述"
// @Param  dict_id body string true "关联字典的id"
// @Param  id  path string true "主键"
// @Router /api/v1/dict-item/{id} [put]
func DictItemEdit(c * gin.Context)  {
	sysDictItem.ID = c.Param("id")
	validate := utils.Verify(&sysDictItem, c)
	if validate == nil {
		data, result := ServiceSys.DictItemUpdate(sysDictItem)
		utils.R(data, result, c)
	}
}
// @Tags 字典
// @Description 字典项删除
// @Param id path string true "主键"
// Router /api/v1/dict-item/{id} [delete]
func DictItemDel(c *gin.Context)  {
	sysDictItem.ID = c.Param("id")
	data, result := ServiceSys.DictItemDel(sysDictItem)
	utils.R(data, result, c)
}