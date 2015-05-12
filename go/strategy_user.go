package client

import (
	"time"
)

type StrategyUser struct {
	Id          int64
	StrategyId  int64
	StrategyOid int64
	Uid         int64
	Expire      time.Time
	IsOpen      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type StrategyUserService struct {

	// 取角色列表 带缓存
	StrategyUserGetListCache func(where Wherer, limit Limiter, sort Sorter) ([]StrategyUser, error)

	// 取角色列表
	StrategyUserGetList func(where Wherer, limit Limiter, sort Sorter) ([]StrategyUser, error)

	// 取角色信息
	StrategyUserGetOneCache func(where Wherer) (*StrategyUser, error)

	// 取角色信息
	StrategyUserGetOne func(where Wherer) (*StrategyUser, error)

	// 创建角色信息
	StrategyUserCreate func(kt_roles *StrategyUser) (int64, error)

	// 删除角色 单条
	StrategyUserDelete func(where Wherer) (int64, error)

	// 更新数据
	StrategyUserUpdate func(id int64, updater map[string]interface{}) (int64, error)

	// 消除缓存
	StrategyUserFlushCache func() error

	// 取角色信息
	StrategyUserGet func(id int64) (*StrategyUser, error)

	// 如果不存在创建
	StrategyUserCreateIfNotExists func(find Wherer, createData *StrategyUser) (int64, error)
}

var StrategyUserClient *StrategyUserService
