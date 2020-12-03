package routerWhiteRouter

import (
	"github.com/gin-gonic/gin"
	ControllerSys "zhoukai/src/controller/sys"
)

func WhiteRouter(e *gin.Engine)  {
	r := e.Group("/api/v1/other/")
	{
		r.GET("/test", ControllerSys.Test)
	}
}