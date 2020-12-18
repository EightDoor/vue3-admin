package model

import (
	"zhoukai/configure"
)

type ResponseData struct {
	// 分页显示，页码
	Page int 	`json:"page"`
	// 分页显示,每页数量
	PageSize int `json:"pageSize"`
	// 返回的元素总数
	//Size int	`json:"size"`
	// 筛选条件计算得到的总数，但可能不会全部返回
	//Total int	`json:"total"`
	List interface{}	`json:"list"`
}


type Response struct {
	// 响应码
	Code configure.Code `json:"code"`
	// 响应描述
	Message interface{}	`json:"message"`
	// 最终返回数据
	Data ResponseData `json:"data"`
}
type Verification struct {
	Message interface{} `json:"message"`
	Code configure.Code `json:"code"`
}

// 公共创建
type CommonCreate struct {
	ID	string `json:"id"`
	CommonTime
}
type CommonTime struct {
	// 默认所有字段的零值, 比如 0, '', false 或者其它 零值，都不会保存到数据库内，使用指针可以避免这种情况。
	CreatedAt int64 `json:"created_at" gorm: "create_at autoCreateTime"`
	UpdatedAt int64 `json:"updated_at" gorm: "update_at autoCreateTime"`
	DeletedAt int64 `json:"deleted_at" gorm:"index autoCreateTime"`
}
