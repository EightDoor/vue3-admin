package router

import (
	"github.com/gin-gonic/gin"
	"zhoukai/controller"
)

func WhiteRouter(e *gin.Engine)  {
	r := e.Group("/api/v1/other/")
	{
		r.POST("/login", controller.UserLogin)
	}
}