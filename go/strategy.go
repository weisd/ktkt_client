package client

import (
	"time"
)

type Strategy struct {
	Id        int64
	Oid       string
	Title     string
	Keywords  string
	Type      int
	DeletedAt time.Time
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"created"`
}

type StrategyService struct {
	StrategyGetListCache func(where Wherer, limit Limiter, sort Sorter) ([]Strategy, error)
}

var StrategyClient *StrategyService
