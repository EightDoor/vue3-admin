package utils

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"zhoukai/src/configure"
	response_data "zhoukai/src/response-data"
)

type Paging struct {
	Page int
	Size int
}
// 公共请求返回
func R(database interface{},result *gorm.DB, c *gin.Context) {
	err := result.Error
	// 总数
	//dates, _ := result.Rows()
	_, pageSize, page := response_data.PageIngService(c)
	responseData := response_data.ResponseData{
		Page:     page,
		PageSize: pageSize,
		List:     database,
	}
	response := response_data.Response{
		Code:    configure.RequestError,
		Message: err,
		Data:    responseData,
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError,response)
	}else {
		response.Code = configure.RequestSuccess
		c.JSON(http.StatusOK, response)
	}

}