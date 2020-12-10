package utils

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"zhoukai/configure"
	"zhoukai/model"
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
	_, pageSize, page := PageIngService(c)
	responseData := model.ResponseData{
		Page:     page,
		PageSize: pageSize,
		List:     database,
	}
	response := model.Response{
		Code:    configure.RequestOtherError,
		Message: err,
		Data:    responseData,
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError,response)
	}else {
		response.Code = configure.RequestSuccess
		response.Message = "success"
		c.JSON(http.StatusOK, response)
	}
}
// 参数校验返回
func Verify(data interface{}, c *gin.Context)  {
	err := c.ShouldBind(data)
	if err != nil {
		response := model.Verification{
			Message: err,
			Code: configure.RequestKeyNotFound,
		}
		c.JSON(http.StatusBadRequest, response)
	}
	//validate := validator.New()
	//err := validate.Struct(modelData)
	//if err == nil {
	//	response := model.Verification{
	//		Message: err,
	//		Code: configure.RequestKeyNotFound,
	//	}
	//	c.JSON(http.StatusBadRequest, response)
	//}
}