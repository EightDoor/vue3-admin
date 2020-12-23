package ServiceSys

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"zhoukai/db"
	ModelSys "zhoukai/model/sys"
	UtilsDB "zhoukai/utils/db"
)

func MenuList(datas []ModelSys.SysMenu, c *gin.Context)([]ModelSys.SysMenu, *gorm.DB)  {
	result := db.DB.Scopes(UtilsDB.Paginate(c, &datas)).Find(&datas)
	return datas, result
}

func MenuCreate(data ModelSys.SysMenu)(ModelSys.SysMenu, *gorm.DB)  {
	data.ID = UtilsDB.CreateUUId()
	result := db.DB.Create(&data)
	return data, result
}
func MenuUpdate(data ModelSys.SysMenu)(ModelSys.SysMenu, *gorm.DB)  {
	result := db.DB.Updates(&data)
	return data, result
}
func MenuDel(data ModelSys.SysMenu)(ModelSys.SysMenu, *gorm.DB)  {
	var sysMenus[] ModelSys.SysMenu
	child := db.DB.Where("dept_id = ?", data.ID).Find(&sysMenus)
	if child.RowsAffected > 0 {
		db.DB.Delete(&sysMenus)
	}
	result := db.DB.Delete(&data)
	return data, result
}
