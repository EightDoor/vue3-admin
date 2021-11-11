package ServiceSys

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"zhoukai/db"
	ModelSys "zhoukai/model/sys"
	UtilsDB "zhoukai/utils/db"
)
func DepartList(list []ModelSys.SysDept, c *gin.Context) ([]ModelSys.SysDept, *gorm.DB)  {
	result := db.DB.Scopes(UtilsDB.Paginate(c, list)).Find(&list)
	return list, result
}
func DepartCreate(data ModelSys.SysDept)(ModelSys.SysDept, *gorm.DB)  {
	data.ID = UtilsDB.CreateUUId()
	result := db.DB.Create(&data)
	return data, result
}
func DepartUpdate(data ModelSys.SysDept)(ModelSys.SysDept, *gorm.DB)  {
	result := db.DB.Updates(data)
	return data, result
}
func DepartDel(data ModelSys.SysDept)(ModelSys.SysDept, *gorm.DB)  {
	var serviceModelDepart []ModelSys.SysDept
	seeChildren := db.DB.Where("dept_id = ?", data.ID).Find(&serviceModelDepart)
	if seeChildren.RowsAffected > 0 {
		db.DB.Delete(&serviceModelDepart)
	}
	result := db.DB.Delete(&data)
	return data, result
}