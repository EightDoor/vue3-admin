package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "gorm.io/gorm/utils"
	"net/http"
	"time"
	"zhoukai/src/configure"
	routerInit "zhoukai/src/router"
)

func main() {
	// 初始化数据库
	configure.InitMySqlCon()
	// 初始化Gin实例
	router := gin.Default()
	routerInit.InitRouter(router)
	// 白名单
	routerInit.WhiteRouter(router)

	// 监听8081端口
	s := &http.Server{
		Addr: ":8081",
		Handler: router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}
}
