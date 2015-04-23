package client

import (
	"time"
)

// 角色表
type KtRoles struct {
	Id       int64
	Title    string
	Keywords string

	DeletedAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type KtRolesService struct {

	// 取角色列表 带缓存
	KtRolesGetListCache func(where Wherer, limit Limiter, sort Sorter) ([]KtRoles, error)

	// 取角色列表
	KtRolesGetList func(where Wherer, limit Limiter, sort Sorter) ([]KtRoles, error)

	// 取角色信息
	KtRolesGetOneCache func(where Wherer) (*KtRoles, error)

	// 取角色信息
	KtRolesGetOne func(where Wherer) (*KtRoles, error)

	// 创建角色信息
	KtRolesCreate func(kt_roles *KtRoles) (int64, error)

	// 删除角色 单条
	KtRolesDelete func(where Wherer) (int64, error)

	// 更新数据
	KtRolesUpdate func(id int64, updater map[string]interface{}) (int64, error)

	// 消除缓存
	KtRolesFlushCache func() error

	// 通过id取info
	KtRolesGet func(id int64) (*KtRoles, error)

	// 通过keywords查找
	KtRolesGetByKeywords func(keywords string) (*KtRoles, error)
}

var KtRolesClient *KtRolesService
