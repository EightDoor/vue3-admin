package UtilsDB

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"zhoukai/model"
)

func PageIngService(c *gin.Context)( pageSize int, page int)   {
	page, _ = strconv.Atoi(c.Query("page"))
	size, _:= strconv.Atoi(c.Query("page_size"))
	if page == 0 {
		page = 1
	}
	if size == 0 {
		size = 10
	}
	return size, page
}

// 分页db统一返回数据
func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB  {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(c.Query("page"))
		if page == 0 {
			page = 1
		}
		pageSize, _ := strconv.Atoi(c.Query("page_size"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize < 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

// 创建添加uuid
func CreateUUId() func(db *gorm.DB) *gorm.DB  {
	return func(db *gorm.DB) *gorm.DB {
		var commonCreate model.CommonCreate
		u2 := uuid.NewV4()
		commonCreate.ID = strings.ReplaceAll(u2.String(), "-", "")
		db.Where("id = ?", 23)
		return db
	}
}
