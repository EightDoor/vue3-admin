package routerV1

import (
	"github.com/gin-gonic/gin"
	ControllerSys "zhoukai/src/controller/sys"
)

func LoadSys(e *gin.RouterGroup)  {
	{
		e.GET("/user", ControllerSys.UserList)
	}
}