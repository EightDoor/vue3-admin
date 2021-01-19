package ModelSys

import (
	_ "github.com/go-playground/validator/v10"
	"time"
	"zhoukai/model"
)

// 部门
type SysDept struct {
	model.CommonCreate
	ParentId string    `json:"parent_id"`
	Name     string `json:"name"`
	OrderNum int16    `json:"order_num"`
}
// 字典
type SysDict struct {
	model.CommonCreate
	// 名称
	Name   string `json:"name" binding:"required"`
	// 编号
	SerialNumber  string `json:"serial_number"`
	// 描述
	Describe	string `json:"describe"`
 }
// 字典项
type SysDictItem struct {
	model.CommonCreate
	// 数据值
	Value	string `json:"value"`
	// 名称
	Label	string `json:"label"`
	// 描述
	Describe string `json:"describe"`
	// 关联的子项item
	DictItem	string `json:"dict_item"`
}

type SysMenu struct {
	model.CommonCreate
	ParentId string    `json:"parent_id" binding:"required"`
	// 菜单名称
	Title     string `json:"title" binding:"required"`
	// 菜单类型： 1. 目录 2. 菜单  3. 按钮
	Type     int    `json:"type" binding:"required"`
	// 排序
	OrderNum int    `json:"order_num" binding:"required"`
	// 是否首页
	IsHome bool	`json:"is_home" binding: "required"`
	// 权限标识，接口标识
	Perms    string `json:"perms"`
	// 菜单标识，前端路由name
	Name     string `json:"name"`
	// 路由路径 (第一级带/ 子级没有/)
	Path	string	`json:"path"`
	// 组件地址
	Component	string	`json:"component"`
	// 重定向地址
	Redirect	string	`json:"redirect"`
	// 图标
	Icon	string	`json:"icon"`
	// 左侧菜单是否隐藏
	Hidden	int	`json:"hidden"`
}

type SysOss struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	Url        string    `xorm:"not null comment('图片等url链接') VARCHAR(255)"`
	Location   string    `xorm:"not null comment('文件存放位置，暂存本地服务器，不采用云 oss') VARCHAR(100)"`
	CreateDate time.Time `xorm:"not null default 'CURRENT_TIMESTAMP(6)' comment('创建时间') TIMESTAMP(6)"`
	Type       string    `xorm:"not null comment('文件类型') VARCHAR(20)"`
	Size       int       `xorm:"not null comment('文件大小 size') INT(11)"`
	OldName    string    `xorm:"not null comment('原文件名称') VARCHAR(100)"`
}

type SysRole struct {
	model.CommonCreate
	Remark     string    `json:"remark" binding:"required"`
	RoleName   string    `json:"role_name"`
}

type SysRoleDept struct {
	Id     string `json:"id"`
	RoleId string `json:"role_id" binding: "required"`
	DeptId string `json:"dept_id" binding: "required"`
}

type SysRoleMenu struct {
	model.CommonCreate
	ID     string `json:"id" gorm:"primary_key"`
	RoleId string `json:"role_id"`
	MenuId string `json:"menu_id"`
}

// 用户
type SysUser struct {
	model.CommonCreate
	PassWord   string    `json:"pass_word"`
	Account    string    `json:"account"`
	NickName   string    `json:"nick_name"`
	Email      string    `json:"email"`
	Status     int       `json:"status"`
	Avatar     string    `json:"avatar"`
	DeptId     string    `json:"dept_id"`
	PhoneNum   string    `json:"phone_num"`
	DepartInfo SysDept	 `json:"depart_info" gorm:"foreignKey:DeptId"`
}
// 用户拥有的信息，拥有角色，拥有菜单
type SysUserWithTheMenu struct {
	UserInfo SysUser `json:"user_info"`
	Menu interface{} `json:"menu"`
}

type SysUserRole struct {
	model.CommonCreate
	ID string `json:"id"`
	UserId string `json:"user_id"`
	RoleId string `json:"role_id"`
}
type SysUserLogin struct {
	ID string `json:"id"`
	Account string `json:"account" binding: "required"`
	PassWord string `json:"pass_word" binding: "required"`
}
