package client

import (
	"time"
)

type StrategyType int

const (
	STRATEGY_TYPE_STRONG StrategyType = 1
	STRATEGY_TYPE_WEAK   StrategyType = 2
)

type Strategy struct {
	Id        int64        `json:"id"`
	Oid       int64        `json:"oid"`
	Title     string       `json:"title"`
	Keywords  string       `json:"keywords"`
	Type      StrategyType `json:"type"`
	DeletedAt time.Time    `xorm:"deleted" json:"-"`
	CreatedAt time.Time    `xorm:"created" json:"createdAt"`
	UpdatedAt time.Time    `xorm:"updated" json:"updatedAt"`
}

type StrategyService struct {

	// 取角色列表 带缓存
	StrategyGetListCache func(where Wherer, limit Limiter, sort Sorter) ([]Strategy, error)

	// 取角色列表
	StrategyGetList func(where Wherer, limit Limiter, sort Sorter) ([]Strategy, error)

	// 取角色信息
	StrategyGetOneCache func(where Wherer) (*Strategy, error)

	// 取角色信息
	StrategyGetOne func(where Wherer) (*Strategy, error)

	// 创建角色信息
	StrategyCreate func(kt_roles *Strategy) (int64, error)

	// 删除角色 单条
	StrategyDelete func(where Wherer) (int64, error)

	// 更新数据
	StrategyUpdate func(id int64, updater map[string]interface{}) (int64, error)

	// 消除缓存
	StrategyFlushCache func() error

	// 通过title取策略
	StrategyGetByTitle func(title string) (*Strategy, error)
}

var StrategyClient *StrategyService
