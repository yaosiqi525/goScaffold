package controllers

const (
	ERROR             = -1 // 未知错误
	SUCCESS           = 0
	PARAM_ERROR       = 100001 // 参数错误
	PAGE_NOT_FOUND    = 100002 // API 找不到
	ADMIN_NOT_FOUND   = 100003 // 用户信息找不到
	ADMIN_TOKEN_LOSER = 100004 // 用户登录信息丢失,需要从新查找
	PERMISSIONS_ERROR = 100005 // 接口权限不够
)
