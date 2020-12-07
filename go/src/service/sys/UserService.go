package ServiceSys

import (
	"zhoukai/src/configure"
	"zhoukai/src/model"
	"zhoukai/src/utils"
)

// 查询用户列表
func UserQuery(dates []model.SysUser, page utils.Paging)([]model.SysUser, int64, error)  {
	pageSize := 10
	offset := (page.Page - 1)*pageSize
	result := configure.DB.Order("id desc").Offset(offset).Limit(pageSize).Find(&dates)
	return dates, result.RowsAffected, result.Error
}

