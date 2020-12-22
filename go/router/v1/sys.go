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
		e.POST("/depart",ControllerSys.DepartCreate)
		e.PUT("/depart/:id", ControllerSys.DepartUpdate)
		e.DELETE("/depart/:id", ControllerSys.DepartDel)
		// 角色
		e.GET("/role", ControllerSys.RoleList)
		e.POST("/role", ControllerSys.RoleCreate)
		e.PUT("/role/:id", ControllerSys.RoleUpdate)
		e.DELETE("/role/:id", ControllerSys.RoleDel)
		// 菜单
		e.GET("/menu", ControllerSys.MenuList)
		e.POST("/menu", ControllerSys.MenuCreate)
	}
}