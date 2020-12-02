package main

import (
	"github.com/gin-gonic/gin"
	md "zhoukai/middleware"
)

func main() {
	r := gin.Default()

	// 初始化Gin实例
	router := gin.Default()
	v1 := router.Group("/apis/v1/")
	{
		v1.GET("/test")
	}
	// secure v1
	sv1 := router.Group("/apis/v1/auth/")
	sv1.Use(md.JWTAuth())
	{
		sv1.GET("/time", con)
	}
	// 监听8080端口
	r.Run()
}