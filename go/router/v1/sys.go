package routerV1

import (
	"github.com/gin-gonic/gin"
	ControllerSys "zhoukai/controller/sys"
)

func LoadSys(e *gin.RouterGroup)  {
	{
		e.GET("/user", ControllerSys.UserList)
		e.GET("/user/:id", ControllerSys.UserSinge)
		e.POST("/user", ControllerSys.UserCreate)
	}
}