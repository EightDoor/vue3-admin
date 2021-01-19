package routerV1

import (
	"github.com/gin-gonic/gin"
	ControllerSys "zhoukai/controller/sys"
)

func LoadSys(e *gin.RouterGroup)  {
	{
		// 用户
		e.GET("/user", ControllerSys.UserList)
		e.PUT("/user/:id", ControllerSys.UserUpdate)
		e.POST("/user", ControllerSys.UserCreate)
		e.DELETE("/user/:id", ControllerSys.UserDel)
		e.GET("/user/getInfo", ControllerSys.GetUserInfo)
		e.GET("/user/permissions/:id", ControllerSys.UserPermissions)
		e.POST("/user/permissions", ControllerSys.UserAssociatedMenu)
		e.GET("/user/getWithTheMenu", ControllerSys.UserGetWithMenu)
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
		e.GET("/role/permissions/:id", ControllerSys.RolePermissions)
		e.POST("/role/permissions", ControllerSys.RoleAssociatedMenu)
		// 菜单
		e.GET("/menu", ControllerSys.MenuList)
		e.POST("/menu", ControllerSys.MenuCreate)
		e.PUT("/menu/:id", ControllerSys.MenuUpdate)
		e.DELETE("/menu/:id", ControllerSys.MenuDel)
		// 字典
		e.POST("/dict", ControllerSys.DictCreate)
		e.GET("/dict", ControllerSys.DictList)
		e.PUT("/dict/:id", ControllerSys.DictEdit)
		e.DELETE("/dict/:id", ControllerSys.DictDel)
	}
}