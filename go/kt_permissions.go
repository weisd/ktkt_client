package client

import (
	"time"
)

// 权限类型
type NodeType int

const (
	// 策略权限
	PERMISSION_TYPE_STRATEGY NodeType = 1
)

// 角色表
type KtPermissions struct {
	Id       int64
	Title    string
	Keywords string
	NodeType NodeType
	NodeId   int64

	DeletedAt time.Time `xorm:"deleted"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}

type KtPermissionsService struct {

	// 取角色列表 带缓存
	KtPermissionsGetListCache func(where Wherer, limit Limiter, sort Sorter) ([]KtPermissions, error)

	// 取角色列表
	KtPermissionsGetList func(where Wherer, limit Limiter, sort Sorter) ([]KtPermissions, error)

	// 取角色信息
	KtPermissionsGetOneCache func(where Wherer) (*KtPermissions, error)

	// 取角色信息
	KtPermissionsGetOne func(where Wherer) (*KtPermissions, error)

	// 创建角色信息
	KtPermissionsCreate func(kt_roles *KtPermissions) (int64, error)

	// 删除角色 单条
	KtPermissionsDelete func(where Wherer) (int64, error)

	// 更新数据
	KtPermissionsUpdate func(id int64, updater map[string]interface{}) (int64, error)

	// 通过id取info
	KtPermissionsGet func(id int64) (*KtPermissions, error)

	// 消除缓存
	KtPermissionsFlushCache func() error
}

var KtPermissionsClient *KtPermissionsService
