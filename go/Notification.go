package client

import (
	"time"
)

type NotificationCategory int

const (
	MOTIFICATION_CATE_UNKNOWN NotificationCategory = iota
	MOTIFICATION_SYSTEM
	MOTIFICATION_LIVE
)

type Notification struct {
	Id         int64
	CategoryId NotificationCategory
	Content    string

	Resource string
	ResId    int64

	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}

type NotificationService struct {

	// 取角色列表 带缓存
	NotificationGetListCache func(where Wherer, limit Limiter, sort Sorter) ([]Notification, error)

	// 取角色列表
	NotificationGetList func(where Wherer, limit Limiter, sort Sorter) ([]Notification, error)

	// 取角色信息
	NotificationGetOneCache func(where Wherer) (*Notification, error)

	// 取角色信息
	NotificationGetOne func(where Wherer) (*Notification, error)

	// 创建角色信息
	NotificationCreate func(info *Notification) (int64, error)

	// 删除角色 单条
	NotificationDelete func(where Wherer) (int64, error)

	// 更新数据
	NotificationUpdate func(id int64, updater map[string]interface{}) (int64, error)

	// 消除缓存
	NotificationFlushCache func() error

	// 取信息
	NotificationGet func(id int64) (*Notification, error)

	// 如果不存在创建
	NotificationCreateIfNotExists func(find Wherer, createData *Notification) (int64, error)

	// 添加新消息
	NotificationAdd func(info *Notification) (int64, error)
}

var NotificationClient *NotificationService
