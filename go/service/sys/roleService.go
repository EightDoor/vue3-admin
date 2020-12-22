package ServiceSys

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
