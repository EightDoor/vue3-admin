package main

import (
	"github.com/gin-gonic/gin"
	_ "gorm.io/gorm/utils"
	"zhoukai/src/configure"
	"zhoukai/src/middleware"
	routerInit "zhoukai/src/router"
)

func main() {
	// 初始化数据库
	configure.InitMySqlCon()
	// 初始化Gin实例
	router := gin.Default()
	// 请求分页中间件
	router.Use(middleware.Paging())
	// 鉴权路由
	routerInit.InitRouter(router)
	// 路由白名单
	routerInit.WhiteRouter(router)

	// 监听8081端口
	router.Run(":8081")
	//s := &http.Server{
	//	Addr: ":8081",
	//	Handler: router,
	//	ReadTimeout:    10 * time.Second,
	//	WriteTimeout:   10 * time.Second,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//err := s.ListenAndServe()
	//if err != nil {
	//	fmt.Printf("startup service failed, err:%v\n\n", err)
	//}
}
