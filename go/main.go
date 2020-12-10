package main

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	_ "gorm.io/gorm/utils"
	"time"
	"zhoukai/configure"
	"zhoukai/middleware"
	routerInit "zhoukai/router"
)

func init()  {
	// 初始化日志库
	middleware.SetLogs(zap.DebugLevel, configure.LOGFORMAT_CONSOLE, "./log/zap.log")
	// 初始化数据库
	configure.InitMySqlCon()
}

func main() {
	// 初始化Gin实例
	router := gin.New()
	// 加载日志中间zap
	router.Use(ginzap.Ginzap(zap.L(), time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(zap.L(), true))
	// 请求分页中间件
	router.Use(middleware.Paging())
	// 鉴权路由
	routerInit.InitRouter(router)
	// 路由白名单
	routerInit.WhiteRouter(router)

	// 生成环境设置为release
	gin.SetMode(gin.DebugMode)
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
