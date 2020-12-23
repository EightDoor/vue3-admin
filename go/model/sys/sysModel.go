package ModelSys

import (
	_ "github.com/go-playground/validator/v10"
	"time"
	"zhoukai/model"
)

// 部门
type SysDept struct {
	model.CommonCreate
	ParentId string    `json:"parent_id"  binding:"required"`
	Name     string `json:"name" binding:"required"`
	OrderNum int16    `json:"order_num" binding:"required"`
}

type SysMenu struct {
	model.CommonCreate
	ParentId string    `json:"parent_id" binding:"required"`
	// 菜单名称
	Name     string `json:"name" binding:"required"`
	// 菜单类型： 1. 目录 2. 菜单  3. 按钮
	Type     int    `json:"type" binding:"required"`
	// 排序
	OrderNum int    `json:"order_num" binding:"required"`
	// 权限标识，接口标识
	Perms    string `json:"perms"`
	// 菜单标识，前端路由name
	Code     string `json:"code" binding:"required"`
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
	Id     int `xorm:"not null pk autoincr INT(11)"`
	RoleId int `xorm:"comment('角色ID') INT(11)"`
	DeptId int `xorm:"comment('部门ID') INT(11)"`
}

type SysRoleMenu struct {
	Id     int `xorm:"not null pk autoincr INT(11)"`
	RoleId int `xorm:"not null index INT(11)"`
	MenuId int `xorm:"not null index INT(11)"`
}

// 用户
type SysUser struct {
	model.CommonCreate
	PassWord   string    `json:"pass_word" xorm:"comment('用户登录密码') VARCHAR(200)"`
	Account    string    `json:"account" xorm:"not null comment('用户登录账号') VARCHAR(32)" binding:"required"`
	NickName   string    `json:"nick_name" xorm:"comment('用户显示昵称') VARCHAR(32)" binding:"required"`
	Email      string    `json:"email" xorm:"not null comment('邮箱地址') VARCHAR(200)"`
	Status     int       `json:"status" xorm:"not null default 1 comment('所属状态是否有效  1是有效 0是失效') TINYINT(4)" binding:"required"`
	Avatar     string    `json:"avatar" xorm:"not null comment('头像地址') VARCHAR(200)"`
	DeptId     string    `json:"dept_id"`
	PhoneNum   string    `json:"phone_num" xorm:"comment('用户手机号码') VARCHAR(20)"`
	DepartInfo SysDept	 `json:"depart_info" gorm:"foreignKey:DeptId"`
}

type SysUserRole struct {
	Id     int `xorm:"not null pk autoincr INT(11)"`
	UserId int `xorm:"not null index INT(11)"`
	RoleId int `xorm:"not null index INT(11)"`
}
