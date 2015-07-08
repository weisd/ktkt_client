package client

import (
	"time"
)

type StrategyInfo struct {
	Id            int64     `json:"id"`
	StrategyId    string    `json:"strategy_id" xorm:"unique"`
	StopProfit    string    `json:"stop_profit"`
	StopLoss      string    `json:"stop_loss"`
	PercentProfit string    `json:"percent_profit"`
	NumProfit     string    `json:"num_profit"`
	TotalTrades   string    `json:"total_trades"`
	DataType      string    `json:"data_type"`
	TestPlatform  string    `json:"test_platform"`
	TestTime      string    `json:"test_time"`
	Intro         string    `json:"intro" xorm:"TEXT"`
	Hot           string    `json:"hot"`
	Publisher     string    `json:"publisher"`
	DeletedAt     time.Time `json:"deleted_at"`
}

type StrategyInfoService struct {

	// 取角色列表 带缓存
	StrategyInfoGetListCache func(where Wherer, limit Limiter, sort Sorter) ([]StrategyInfo, error)

	// 取角色列表
	StrategyInfoGetList func(where Wherer, limit Limiter, sort Sorter) ([]StrategyInfo, error)
	// 取角色信息
	StrategyInfoGetOneCache func(where Wherer) (*StrategyInfo, error)
	// 取角色信息
	StrategyInfoGetOne func(where Wherer) (*StrategyInfo, error)

	// 删除角色 单条
	StrategyInfoDelete func(where Wherer) (int64, error)

	// 更新数据
	StrategyInfoUpdate func(id int64, updater map[string]interface{}) (int64, error)

	// 消除缓存
	StrategyInfoFlushCache func() error
	// 通过id取info
	StrategyInfoGet func(id int64) (*StrategyInfo, error)

	StrategyInfoUpdateByStrategyId func(strategyId int64, updater map[string]interface{}) (int64, error)

	StrategyInfoGetByStrategyId func(strategyId int64) (*StrategyInfo, error)

	// 创建角色信息
	StrategyInfoCreate func(kt_roles *StrategyInfo) (int64, error)
}

var StrategyInfoClient *StrategyInfoService
