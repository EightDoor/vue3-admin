package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

// 分页统一返回数据
func PageIngService(c *gin.Context)(offset int, pageSize int, page int)   {
	page, _ = strconv.Atoi(c.Query("page"))
	size, _:= strconv.Atoi(c.Query("size"))
	if page == 0 {
		page = 1
	}
	if size == 0 {
		size = 20
	}
	offset = (page - 1) * size
	return offset, size, page
}