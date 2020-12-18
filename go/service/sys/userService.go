package ServiceSys

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"zhoukai/db"
	"zhoukai/model/sys"
	"zhoukai/utils"
	"zhoukai/utils/db"
)

func UserQuery(dates []ModelSys.SysUser, c *gin.Context)([]ModelSys.SysUser, *gorm.DB)  {
	result := db.DB.Scopes(UtilsDB.Paginate(c)).Find(&dates)
	return dates, result
}
func UserSinge(id string)(ModelSys.SysUser, *gorm.DB)  {
	var data ModelSys.SysUser
	result := db.DB.First(&data, id)
	return data, result
}
func UserCreate(data ModelSys.SysUser)(ModelSys.SysUser, *gorm.DB)  {
	data.ID = UtilsDB.CreateUUId()
	data.PassWord = utils.GetMd5String(data.PassWord)
	result := db.DB.Create(&data)
	return data, result
}
func UserUpdate(data ModelSys.SysUser)(ModelSys.SysUser, *gorm.DB)  {
	var result *gorm.DB
	if data.PassWord != "" {
		passwdR := db.DB.Find(&data, data.ID)
		data.PassWord = utils.GetMd5String(data.PassWord)
		if passwdR.Error != nil && passwdR.RowsAffected != 0 {
			result = db.DB.Updates(&data)
		}
	}
	result = db.DB.Omit("pass_word").Updates(&data)
	return data, result
}
func UserDel(data ModelSys.SysUser) *gorm.DB  {
	result := db.DB.Delete(&data)
	return result
}


