package routerWhiteRouter

import (
	"zhoukai/src/configure"
)

type ResponseData struct {
	// 分页显示，页码
	Page int64 	`json:"page"`
	// 分页显示,每页数量
	PageSize int64 `json:"page_size"`
	// 返回的元素总数
	Size int64	`json:"size"`
	// 筛选条件计算得到的总数，但可能不会全部返回
	Total int64	`json:"total"`
	List []interface{}	`json:"list"`
}

type Response struct {
	// 响应码
	Code configure.Code `json:"code"`
	// 响应描述
	Message string 		`json:"message"`
	// 最终返回数据
	Data ResponseData `json:"data"`
}