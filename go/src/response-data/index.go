package response_data

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"zhoukai/src/configure"
)

// 分页统一返回数据
func PageIngService(c *gin.Context)(offset int, pageSize int, page int)   {
	page, _ = strconv.Atoi(c.Query("page"))
	size, _:= strconv.Atoi(c.Query("size"))
	offset = (page - 1) * size
	return offset, size, page
}

type ResponseData struct {
	// 分页显示，页码
	Page int 	`json:"page"`
	// 分页显示,每页数量
	PageSize int `json:"pageSize"`
	// 返回的元素总数
	//Size int	`json:"size"`
	// 筛选条件计算得到的总数，但可能不会全部返回
	//Total int	`json:"total"`
	List interface{}	`json:"list"`
}


type Response struct {
	// 响应码
	Code configure.Code `json:"code"`
	// 响应描述
	Message error		`json:"message"`
	// 最终返回数据
	Data ResponseData `json:"data"`
}