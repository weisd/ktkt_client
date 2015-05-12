package client

// 角色权限
type KtRolePermission struct {
	Id           int64
	RolesId      int64
	PermissionId int64
}

type KtRolePermissionService struct {

	// 取角色列表 带缓存
	KtRolePermissionGetListCache func(where Wherer, limit Limiter, sort Sorter) ([]KtRolePermission, error)

	// 取角色列表
	KtRolePermissionGetList func(where Wherer, limit Limiter, sort Sorter) ([]KtRolePermission, error)

	// 取角色信息
	KtRolePermissionGetOneCache func(where Wherer) (*KtRolePermission, error)

	// 取角色信息
	KtRolePermissionGetOne func(where Wherer) (*KtRolePermission, error)

	// 创建角色信息
	KtRolePermissionCreate func(kt_roles *KtRolePermission) (int64, error)

	// 删除角色 单条
	KtRolePermissionDelete func(where Wherer) (int64, error)

	// 更新数据
	KtRolePermissionUpdate func(id int64, updater map[string]interface{}) (int64, error)

	// 消除缓存
	KtRolePermissionFlushCache func() error

	// 如果不存在创建
	KtRolePermissionCreateIfNotExists func(find Wherer, createData *KtRolePermission) (int64, error)
}

var KtRolePermissionClient *KtRolePermissionService
