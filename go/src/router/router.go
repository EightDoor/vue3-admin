package routerWhiteRouter

import (
	"github.com/gin-gonic/gin"
	routerV1 "zhoukai/src/router/v1"
	md "zhoukai/src/middleware"
)

func InitRouter(e *gin.Engine)  {
	v1 := e.Group("/api/v1/")
	// 鉴权路由
	v1.Use(md.JWTAuth())
	// sys
	routerV1.LoadSys(v1)
}