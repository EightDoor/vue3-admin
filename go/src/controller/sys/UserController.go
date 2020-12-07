package ControllerSys

import (
	"github.com/gin-gonic/gin"
	http "net/http"
	"strconv"
	"zhoukai/src/model"
	ServiceSys "zhoukai/src/service/sys"
	"zhoukai/src/utils"
)

func UserList(c *gin.Context)  {
	page, _ := strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = 1
	}
	paging := utils.Paging{
		Page: page,
		Size: 10,
	}
	var list []model.SysUser
	res, _ , err := ServiceSys.UserQuery(list, paging)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":  err,
			"data": res,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":  "",
		"data": res,
	})
}