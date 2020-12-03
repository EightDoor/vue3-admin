package routerV1

import "github.com/gin-gonic/gin"

func LoadSys(e *gin.RouterGroup)  {
	{
		e.GET("/test")
	}
}