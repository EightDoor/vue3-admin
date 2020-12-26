package ServiceSys

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
	"zhoukai/db"
	"zhoukai/model/sys"
	"zhoukai/utils"
	"zhoukai/utils/db"
)

func UserQuery(dates []ModelSys.SysUser, c *gin.Context)([]ModelSys.SysUser, *gorm.DB) {
	result :=db.DB.Scopes(UtilsDB.Paginate(c, dates)).Preload("DepartInfo").Find(&dates)
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
func UserGetInfo(data ModelSys.SysUser)(ModelSys.SysUser ,*gorm.DB)  {
	result := db.DB.Preload("DepartInfo").First(&data)
	return data, result
}

func UserPermission(id string, datas []ModelSys.SysRole, c *gin.Context)([]ModelSys.SysRole, *gorm.DB)  {
	var sysUserPermi []ModelSys.SysUserRole
	var list []string
	var count int64
	var result *gorm.DB
	findList := db.DB.Where("user_id = ?", id).Find(&sysUserPermi)
	if findList.RowsAffected > 0 {
		for _,value := range sysUserPermi{
			list = append(list, value.RoleId)
		}
		result = db.DB.Find(&datas, list).Count(&count)
	} else {
		datas = make([]ModelSys.SysRole, 0)
		result = new(gorm.DB)
	}
	UtilsDB.SetTotal(c, count)
	return datas, result
}
func UserAssociatedMenu(data ModelSys.SysUserRole)(ModelSys.SysUserRole, *gorm.DB)  {
	var sysUserMenu ModelSys.SysUserRole
	var datas []ModelSys.SysUserRole
	var permission []ModelSys.SysUserRole
	var result *gorm.DB
	userId := data.UserId
	roleId := strings.Split(data.RoleId, ",")
	// 先删除存在的权限，然后存储权限
	rr := db.DB.Where("user_id = ?", userId).Find(&permission)
	if rr.RowsAffected > 0 {
		db.DB.Delete(&permission)
	}

	for i := 0;i<len(roleId); i++  {
		r := db.DB.Where("role_id = ? AND menu_id = ?", roleId, roleId[i]).Find(&sysUserMenu)
		if r.RowsAffected == 0 {
			datas = append(datas, ModelSys.SysUserRole{
				ID: UtilsDB.CreateUUId(),
				UserId: userId,
				RoleId: roleId[i],
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


