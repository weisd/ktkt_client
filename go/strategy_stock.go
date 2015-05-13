package client

import (
	"time"
)

type StrategyStock struct {
	Id         int64     `json:"id"`
	StrategyId string    `json:"strategy_id"`
	Stkname    string    `json:"stkname"`
	Stkcode    string    `json:"stkcode"`
	CodeNo     string    `json:"code_no"`
	Fmlprice   string    `json:"fmlprice"`
	Fmltime    string    `json:"fmltime"`
	Fmlname    string    `json:"fmlname"`
	Date       string    `json:"date"`
	Up         int       `json:"up"`
	Down       int       `json:"down"`
	CreatedAt  time.Time `xorm:"created" json:"createdAt"`
	UpdatedAt  time.Time `xorm:"updated" json:"updatedAt"`
}

type StrategyStockService struct {

	// 取角色列表 带缓存
	StrategyStockGetListCache func(where Wherer, limit Limiter, sort Sorter) ([]StrategyStock, error)

	// 取角色列表
	StrategyStockGetList func(where Wherer, limit Limiter, sort Sorter) ([]StrategyStock, error)

	// 取角色信息
	StrategyStockGetOneCache func(where Wherer) (*StrategyStock, error)

	// 取角色信息
	StrategyStockGetOne func(where Wherer) (*StrategyStock, error)

	// 创建角色信息
	StrategyStockCreate func(kt_roles *StrategyStock) (int64, error)

	// 删除角色 单条
	StrategyStockDelete func(where Wherer) (int64, error)

	// 更新数据
	StrategyStockUpdate func(id int64, updater map[string]interface{}) (int64, error)

	// 消除缓存
	StrategyStockFlushCache func() error

	// 取信息
	StrategyStockGet func(id int64) (*StrategyStock, error)

	// 如果不存在创建
	StrategyStockCreateIfNotExists func(find Wherer, createData *StrategyStock) (int64, error)
}

var StrategyStockClient *StrategyStockService
