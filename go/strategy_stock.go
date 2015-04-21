package client

import (
	"time"
)

type StrategyStock struct {
	Id         int64
	StrategyId int64
	Stkname    string
	Stkcode    string
	CodeNo     string
	Fmlprice   string
	Fmltime    string
	Fmlname    string
	Date       string
	Up         int
	Down       int
	CreatedAt  time.Time `xorm:"created"`
	UpdatedAt  time.Time `xorm:"created"`
}

type StrategyStockService struct {
	StrategyStockGetListCache func(where Wherer, limit Limiter, sort Sorter) ([]StrategyStock, error)
	StrategyStockCreate       func(stock *StrategyStock) error
}

var StrategyStockClient *StrategyStockService
