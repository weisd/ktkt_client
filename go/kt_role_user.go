package client

// 角色 用户关系表
type KtRoleUser struct {
	Id      int64
	UserId  int64
	RolesId int64
}

type KtRoleUserService struct {

	// 取角色列表 带缓存
	KtRoleUserGetListCache func(where Wherer, limit Limiter, sort Sorter) ([]KtRoleUser, error)

	// 取角色列表
	KtRoleUserGetList func(where Wherer, limit Limiter, sort Sorter) ([]KtRoleUser, error)

	// 取角色信息
	KtRoleUserGetOneCache func(where Wherer) (*KtRoleUser, error)

	// 取角色信息
	KtRoleUserGetOne func(where Wherer) (*KtRoleUser, error)

	// 创建角色信息
	KtRoleUserCreate func(kt_roles *KtRoleUser) (int64, error)

	// 删除角色 单条
	KtRoleUserDelete func(where Wherer) (int64, error)

	// 更新数据
	KtRoleUserUpdate func(id int64, updater map[string]interface{}) (int64, error)

	// 消除缓存
	KtRoleUserFlushCache func() error
}

var KtRoleUserClient *KtRoleUserService
