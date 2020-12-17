package service

// 公共创建
type CommonCreate struct {
	Id	string `json:"id"`
	CommonTime
}
type CommonTime struct {
	CreateTime int16 `json:"create_time"`
	UpdateTime int16 `json:"update_time"`
	DeleteTime int16 `json:"delete_time"`
}
