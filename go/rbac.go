package client

type Rbac struct {
}

type RbacService struct {

	// 角色下的所有权限
	GetRolePermissionIds func(roldId int64) ([]int64, error)

	// 角色下的所有节点
	GetRoleNodeIds func(roldId int64, noteType NodeType) ([]string, error)

	// 取用户对应的角色ids
	GetUserRoleIds func(uid int64) ([]int64, error)

	// 取用户对应的角色
	GetUserRoles func(uid int64) ([]*KtRoles, error)

	// 权限对应的所有角色
	GetPermissionRoleIds func(permissionId int64) ([]int64, error)

	// 权限对应的所有角色
	GetPermissionRoles func(permissionId int64) ([]*KtRoles, error)

	// 通过nodetype取permission
	GetPermissionByNodeId func(nodeId string, nodeType NodeType) (*KtPermissions, error)

	// 权限节点对应的所有角色
	GetRolesByNodeId func(nodeId string, nodeType NodeType) ([]*KtRoles, error)

	// 权限节点对应的所有角色
	GetRoleIdsByNodeId func(nodeId string, nodeType NodeType) ([]int64, error)

	// 用户对应的所有权限id
	GetUserPermissionIds func(uid int64) ([]int64, error)

	// 用户对应的所有权限节点id
	GetUserNodeIds func(uid int64, nodeType NodeType) ([]string, error)

	// 用户是否于角色
	HasRoles func(uid int64, keywords string) (bool, error)

	// 用户是否有node权限
	HasNode func(uid int64, nodeId string, nodeType NodeType) (bool, error)

	// 通过角色id取所有用户id
	GetUserIdsByRoleId func(roleId int64) ([]int64, error)

	// 通过节点id取所有用户id
	GetUserIdsByNodeId func(nodeId string, nodeType NodeType) ([]int64, error)

	// 开启发送功能的用户ids
	GetStockSendUidsByStrategyId func(strategyId string) []int64
	// 删除发送功能
	DelStockSendUserByStrategyId func(strategyId string, uid int64) (affected int64, err error)
	// 删除发送功能
	AddStockSendUserByStrategyId func(strategyId string, uids []int64) (affected int64, err error)
	// 关闭用户的所有
	ClearStockSendByUid func(uid int64)
	// 所有策略
	GetAllStrategyNodeIds func() ([]string, error)
}

var RbacClient *RbacService
