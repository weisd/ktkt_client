package client

import (
	"time"
)

type Strategy struct {
	Id        int64     `json:"id"`
	Oid       string    `json:"oid"`
	Title     string    `json:"title"`
	Keywords  string    `json:"keywords"`
	Type      int       `json:"type"`
	DeletedAt time.Time `xorm:"deleted" json:"_"`
	CreatedAt time.Time `xorm:"created" json:"_"`
	UpdatedAt time.Time `xorm:"updated" json:"_"`
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
}

var StrategyClient *StrategyService
