package client

import (
	"time"
)

type Clients struct {
	Id     int64  `json:"id"`
	Title  string `json:"title"`
	Slug   string `json:"slug"`
	Secret string `json:"secret"`

	CreatedAt time.Time `xorm:"created" json:"created_at"`
	UpdatedAt time.Time `xorm:"updated" json:"updated_at"`
}

type ClientsService struct {

	// 取角色列表 带缓存
	ClientsGetListCache func(where Wherer, limit Limiter, sort Sorter) ([]Clients, error)

	// 取角色列表
	ClientsGetList func(where Wherer, limit Limiter, sort Sorter) ([]Clients, error)

	// 取角色信息
	ClientsGetOneCache func(where Wherer) (*Clients, error)

	// 取角色信息
	ClientsGetOne func(where Wherer) (*Clients, error)

	// 创建角色信息
	ClientsCreate func(kt_roles *Clients) (int64, error)

	// 删除角色 单条
	ClientsDelete func(where Wherer) (int64, error)

	// 更新数据
	ClientsUpdate func(id int64, updater map[string]interface{}) (int64, error)

	// 消除缓存
	ClientsFlushCache func() error

	// 通过title取策略
	ClientsGetByTitle func(title string) (*Clients, error)

	// 取信息
	ClientsGet func(id int64) (*Clients, error)

	// 如果不存在创建
	ClientsCreateIfNotExists func(find Wherer, createData *Clients) (int64, error)
}

var ClientsClient *ClientsService
