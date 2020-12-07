package model

import (
	_ "github.com/go-playground/validator/v10"
	"time"
)

type SysDept struct {
	DeptId   int    `xorm:"not null pk autoincr INT(11)"`
	ParentId int    `xorm:"not null INT(11)"`
	Name     string `xorm:"not null comment('部门名称') VARCHAR(50)"`
	DelFlag  int    `xorm:"not null default 1 TINYINT(4)"`
	OrderNum int    `xorm:"not null comment('排序') INT(11)"`
}

type SysMenu struct {
	MenuId   int    `xorm:"not null pk autoincr INT(11)"`
	ParentId int    `xorm:"not null INT(11)"`
	Name     string `xorm:"not null comment('菜单名称') VARCHAR(50)"`
	Type     int    `xorm:"not null comment('菜单类型： 1. 菜单/目录 2 tabs 3 按钮') INT(11)"`
	OrderNum int    `xorm:"not null comment('排序') INT(11)"`
	Perms    string `xorm:"comment('权限标识，接口标识') VARCHAR(255)"`
	Code     string `xorm:"not null comment('菜单标识，前端路由name') VARCHAR(30)"`
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
	RoleId     int       `xorm:"not null pk autoincr INT(11)"`
	Remark     string    `xorm:"not null comment('角色备注') VARCHAR(100)"`
	RoleName   string    `xorm:"not null comment('角色名称') VARCHAR(100)"`
	DeptId     int       `xorm:"comment('关联部门ID') INT(11)"`
	CreateDate time.Time `xorm:"not null default 'CURRENT_TIMESTAMP(6)' comment('创建时间') TIMESTAMP(6)"`
}

type SysRoleDept struct {
	Id     int `xorm:"not null pk autoincr INT(11)"`
	RoleId int `xorm:"comment('角色ID') INT(11)"`
	DeptId int `xorm:"comment('部门ID') INT(11)"`
}

type SysRoleMenu struct {
	Id     int `xorm:"not null pk autoincr INT(11)"`
	RoleId int `xorm:"not null index INT(11)"`
	MenuId int `xorm:"not null index INT(11)"`
}

type SysUser struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	Password   string    `xorm:"comment('用户登录密码') VARCHAR(200)" validate:"reqired"`
	Account    string    `xorm:"not null comment('用户登录账号') VARCHAR(32)"`
	NickName   string    `xorm:"comment('用户显示昵称') VARCHAR(32)"`
	Email      string    `xorm:"not null comment('邮箱地址') VARCHAR(200)"`
	Status     int       `xorm:"not null default 1 comment('所属状态是否有效  1是有效 0是失效') TINYINT(4)"`
	Avatar     string    `xorm:"not null comment('头像地址') VARCHAR(200)"`
	DeptId     int       `xorm:"not null unique INT(11)"`
	CreateDate time.Time `xorm:"not null default 'CURRENT_TIMESTAMP(6)' comment('创建时间') TIMESTAMP(6)"`
	UpdateDate time.Time `xorm:"not null default 'CURRENT_TIMESTAMP(6)' comment('更新时间') TIMESTAMP(6)"`
	PhoneNum   string    `xorm:"comment('用户手机号码') VARCHAR(20)"`
}

type SysUserRole struct {
	Id     int `xorm:"not null pk autoincr INT(11)"`
	UserId int `xorm:"not null index INT(11)"`
	RoleId int `xorm:"not null index INT(11)"`
}
