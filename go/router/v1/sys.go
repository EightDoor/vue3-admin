package routerV1

import (
	"github.com/gin-gonic/gin"
	ControllerSys "zhoukai/controller/sys"
)

func LoadSys(e *gin.RouterGroup)  {
	{
		// 用户
		e.GET("/user", ControllerSys.UserList)
		e.GET("/user/:id", ControllerSys.UserSinge)
		e.PUT("/user/:id", ControllerSys.UserUpdate)
		e.POST("/user", ControllerSys.UserCreate)
		e.DELETE("/user/:id", ControllerSys.UserDel)
		// 部门
		e.GET("/depart", ControllerSys.DepartList)
	}
}