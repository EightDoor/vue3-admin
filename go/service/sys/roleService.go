package ServiceSys

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"zhoukai/db"
	ModelSys "zhoukai/model/sys"
	UtilsDB "zhoukai/utils/db"
)

func RoleList(datas []ModelSys.SysRole, c *gin.Context)([]ModelSys.SysRole, *gorm.DB)  {
	result := db.DB.Scopes(UtilsDB.Paginate(c, datas)).Find(&datas)
	return datas, result
}

func RoleCreate(data ModelSys.SysRole)(ModelSys.SysRole, *gorm.DB)  {
	data.ID = UtilsDB.CreateUUId()
	result := db.DB.Create(&data)
	return data, result
}
func RoleUpdate(data ModelSys.SysRole)(ModelSys.SysRole, *gorm.DB)  {
	result := db.DB.Updates(&data)
	return data, result
}
func RoleDel(data ModelSys.SysRole)(ModelSys.SysRole, *gorm.DB)  {
	result := db.DB.Delete(&data)
	return data, result
}
func RolePermission(id string, datas []ModelSys.SysMenu, c *gin.Context)([]ModelSys.SysMenu, *gorm.DB)  {
	var sysRolePermi []ModelSys.SysRoleMenu
	var list []int
	var count int64
	db.DB.Where("role_id = ?", id).Find(&sysRolePermi)
	for _,value := range sysRolePermi{
		result, _ := strconv.Atoi(value.MenuId)
		list = append(list, result)
	}
	result := db.DB.Find(&datas, list).Count(&count)
	UtilsDB.SetTotal(c, count)
	return datas, result
}
