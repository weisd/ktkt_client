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
}

var RbacClient *RbacService
