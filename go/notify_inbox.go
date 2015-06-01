package client

import (
	"time"
)

type NotifyInbox struct {
	Id         int64
	Uid        int64
	Mid        int64
	CategoryId NotificationCategory
	Readed     bool `xorm:"default 0"`

	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:""`
}

type NotifyInboxService struct {

	// 取角色列表 带缓存
	NotifyInboxGetListCache func(where Wherer, limit Limiter, sort Sorter) ([]NotifyInbox, error)

	// 取角色列表
	NotifyInboxGetList func(where Wherer, limit Limiter, sort Sorter) ([]NotifyInbox, error)

	// 取角色信息
	NotifyInboxGetOneCache func(where Wherer) (*NotifyInbox, error)

	// 取角色信息
	NotifyInboxGetOne func(where Wherer) (*NotifyInbox, error)

	// 创建角色信息
	NotifyInboxCreate func(info *NotifyInbox) (int64, error)

	// 删除角色 单条
	NotifyInboxDelete func(where Wherer) (int64, error)

	// 更新数据
	NotifyInboxUpdate func(id int64, updater map[string]interface{}) (int64, error)

	// 消除缓存
	NotifyInboxFlushCache func() error

	// 取信息
	NotifyInboxGet func(id int64) (*NotifyInbox, error)

	// 如果不存在创建
	NotifyInboxCreateIfNotExists func(find Wherer, createData *NotifyInbox) (int64, error)

	NotifyInboxGetListByUid func(uid int64, limit Limiter, sort Sorter) ([]*Notification, error)
}

var NotifyInboxClient *NotifyInboxService
