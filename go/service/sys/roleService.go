package ServiceSys

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
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
	var list []string
	var count int64
	var result *gorm.DB
	findList := db.DB.Where("role_id = ?", id).Find(&sysRolePermi)
	if findList.RowsAffected > 0 {
		for _,value := range sysRolePermi{
			list = append(list, value.MenuId)
		}
		result = db.DB.Find(&datas, list).Count(&count)
	}else {
		datas = make([]ModelSys.SysMenu, 0)
		result = new(gorm.DB)
	}
	UtilsDB.SetTotal(c, count)
	return datas, result
}
func RoleAssociatedMenu(data ModelSys.SysRoleMenu)(ModelSys.SysRoleMenu, *gorm.DB)  {
	var sysRoleMenu ModelSys.SysRoleMenu
	var datas []ModelSys.SysRoleMenu
	var permission []ModelSys.SysRoleMenu
	var result *gorm.DB
	roleId := data.RoleId
	menuId := strings.Split(data.MenuId, ",")
	// 先删除存在的权限，然后存储权限
	rr := db.DB.Where("role_id = ?", roleId).Find(&permission)
	if rr.RowsAffected > 0 {
		db.DB.Delete(&permission)
	}

	for i := 0;i<len(menuId); i++  {
		r := db.DB.Where("role_id = ? AND menu_id = ?", roleId, menuId[i]).Find(&sysRoleMenu)
		if r.RowsAffected == 0 {
			datas = append(datas, ModelSys.SysRoleMenu{
				ID: UtilsDB.CreateUUId(),
				RoleId: roleId,
				MenuId: menuId[i],
			})
		}
	}
	if len(datas) > 0 {
		result = db.DB.Create(&datas)
	}else {
		result = new(gorm.DB)
	}
	return data, result
}
