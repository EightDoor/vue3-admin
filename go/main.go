package main

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	_ "gorm.io/gorm/utils"
	"time"
	"zhoukai/configure"
	"zhoukai/db"
	"zhoukai/docs"
	"zhoukai/middleware"
	routerInit "zhoukai/router"
)

func init()  {
	// 初始化日志库
	middleware.SetLogs(zap.DebugLevel, configure.LOGFORMAT_CONSOLE, "./log/zap.log")
	// 初始化数据库
	db.InitMySqlCon()
}


// @title vue3-go-admin 后台管理接口文档
// @version 1.0
// @description 后台管理

// @contact.name 周凯
// @contact.url http://www.start6.cn
// @contact.email zhou851708184@163.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host http://localhost:8081/swagger/index.html
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
	// 加载swagger
	docs.SwaggerInfo.Title = "vue3-go-admin后台接口文档"
	docs.SwaggerInfo.Description = "后台管理"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8081"
	//docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 生产环境设置为release
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
