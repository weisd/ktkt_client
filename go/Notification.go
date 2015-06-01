package client

import (
	"time"
)

type NotificationCategory int

const (
	NOTIFICATION_CATE_UNKNOWN NotificationCategory = iota
	NOTIFICATION_SYSTEM
	NOTIFICATION_LIVE
)

type Notification struct {
	Id         int64                `json:"id"`
	CategoryId NotificationCategory `json:"category_id"`
	Content    string               `json:"content"`

	Resource string `json:"resource"`
	ResId    string `json:"res_id"`

	CreatedAt time.Time `xorm:"created" json:"created_at"`
	UpdatedAt time.Time `xorm:"updated"  json:"updated_at`
	DeletedAt time.Time `xorm:"deleted" json:"deleted_at`
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
