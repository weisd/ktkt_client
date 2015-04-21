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
	StrategyUserGetListCache func(where Wherer, limit Limiter, sort Sorter) ([]StrategyUser, error)
}

var StrategyUserClient *StrategyService
