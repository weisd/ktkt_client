package client

import (
	"time"
)

type Mystock struct {
	Id      int64  `json:"id"`
	Uid     int64  `json:"uid"`
	CodeNo  string `json:"code_no"`
	Code    string `json:"code"`
	GroupId int64  `json:"groupId"`
	Group   string `json:"group"`
	Sort    int64  `json:"sort" xorm:"default 0"`

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

	// 取信息
	MystockGet func(id int64) (*Mystock, error)

	// 删除用户所有自选股
	MystockDeleteByUid func(uid int64) (int64, error)

	// 删除分组下的自选
	MystockDeleteByGroupId func(id int64) (int64, error)

	MystockDeleteByUidAndGroupId func(uid, groupId int64) (int64, error)

	// 如果不存在创建
	MystockCreateIfNotExists func(find Wherer, createData *Mystock) (int64, error)

	// 批量创建
	MystockMultiCreate func(codes []*Mystock) []int64

	// total
	MystockCount func(w Wherer) (int64, error)

	// 多删除
	MystockMultiDeleteByUid func(uid int64, mystocks []*Mystock) (int64, error)

	MystockGroupListWithStatByUid func(uid int64, code string) ([]map[string]interface{}, error)
}

var MystockClient *MystockService
