package middleware

import (
	"github.com/gin-gonic/gin"
)

// 分页
func Paging() gin.HandlerFunc  {
	return func(c *gin.Context) {
		//page, _ := strconv.Atoi(c.Query("page"))
		//size, _:= strconv.Atoi(c.Query("size"))
		//paging := utils.Paging{
		//	Page: page,
		//	Size: size,
		//}
		c.Set("pagingData", "123")
	}
}