package UtilsDB

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"strconv"
	"strings"
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
		// TODO 最大的page_size为1000条
		case pageSize > 1000:
			pageSize = 1000
		case pageSize < 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

// 创建添加uuid
func CreateUUId() string  {
	u2 := uuid.NewV4()
	return strings.ReplaceAll(u2.String(), "-", "")
}

//func CreateUUId() func(db *gorm.DB) *gorm.DB  {
//	return func(db *gorm.DB) *gorm.DB {
//		var commonCreate model.CommonCreate
//		u2 := uuid.NewV4()
//		commonCreate.ID = strings.ReplaceAll(u2.String(), "-", "")
//		db.Where("id = ?", 23)
//		return db
//	}
//}
func SetTotal(c *gin.Context, count int64)  {
	c.Set("TotalCustom", count)
}
func GetTotal(c *gin.Context) interface{}  {
	r, _ := c.Get("TotalCustom")
	return r
}
