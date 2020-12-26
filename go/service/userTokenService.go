package service

import (
	"gorm.io/gorm"
	"zhoukai/db"
	ModelSys "zhoukai/model/sys"
	"zhoukai/utils"
)

func Login(data ModelSys.SysUserLogin)(ModelSys.SysUserLogin, *gorm.DB, int)  {
	var result *gorm.DB
	var sysUserR ModelSys.SysUserLogin
	// 1 是账户名不存在 2是密码错误
	var status int
	findUser := db.DB.Table("sys_user").Where("account = ?", data.Account).Find(&sysUserR)
	if findUser.RowsAffected > 0 {
		data.PassWord = utils.GetMd5String(data.PassWord)
		result = db.DB.Table("sys_user").Where("account = ? AND pass_word = ?", data.Account, data.PassWord).First(&data)
	}else {
		status = 1
		result = new(gorm.DB)
	}
	return data, result, status
}
