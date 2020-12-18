package ServiceSys

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"zhoukai/db"
	ModelSys "zhoukai/model/sys"
	UtilsDB "zhoukai/utils/db"
)
func DepartList(list []ModelSys.SysDept, c *gin.Context) ([]ModelSys.SysDept, *gorm.DB)  {
	result := db.DB.Scopes(UtilsDB.Paginate(c)).Find(&list)
	return list, result
}