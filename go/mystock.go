package client

import (
	"time"
)

type Mystock struct {
	Id     int64  `json:"id"`
	Uid    int64  `json:"uid"`
	CodeNo string `json:"code_no"`
	Code   string `json:"code"`
	Type   int64  `json:"type"`

	DeletedAt time.Time `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type MystockService struct {

	// 取列表 带缓存
	MystockGetListCache func(where Wherer, limit Limiter, sort Sorter) ([]Mystock, error)

	// 取列表
	MystockGetList func(where Wherer, limit Limiter, sort Sorter) ([]Mystock, error)

	// 取信息
	MystockGetOneCache func(where Wherer) (*Mystock, error)

	// 取信息
	MystockGetOne func(where Wherer) (*Mystock, error)

	// 创建信息
	MystockCreate func(kt_roles *Mystock) (int64, error)

	// 删除 单条
	MystockDelete func(where Wherer) (int64, error)

	// 更新数据
	MystockUpdate func(id int64, updater map[string]interface{}) (int64, error)

	// 消除缓存
	MystockFlushCache func() error
}

var MystockClient *MystockService
