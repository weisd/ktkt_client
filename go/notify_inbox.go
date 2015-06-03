package client

import (
	"time"
)

type NotifyInbox struct {
	Id         int64                `json:"id"`
	Uid        int64                `xorm:"index" json:"uid"`
	Mid        int64                `json:"mid"`
	CategoryId NotificationCategory `json:"category_id"`
	Title      string               `json:"title"`
	Content    string               `json:"content"`

	Resource string `json:"resource"`
	ResId    string `json:"res_id"`

	Readed bool `xorm:"default 0" json:"readed"`

	CreatedAt time.Time `xorm:"created" json:"created_at"`
	UpdatedAt time.Time `xorm:"updated"  json:"updated_at"`
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

	NotifyInboxGetListByUid func(uid int64, limit Limiter, sort Sorter) ([]NotifyInbox, error)

	NotifyInboxGetListByUidAnCate func(uid int64, category NotificationCategory, limit Limiter, sort Sorter) ([]NotifyInbox, error)

	NotifyInboxCountByUid func(uid int64) (int64, error)

	NotifyInboxGetUnRead func(uid int64) (map[string]int64, error)

	NotifyInboxUpdateUnRead func(uid int64) error

	NotifyInboxMultiRead func(uid int64, ids []int64) int64

	NotifyInboxMultiDelete func(uid int64, ids []int64) int64

	NotifyInboxGetById func(id int64, withContent bool) (*NotifyInbox, error)
}

var NotifyInboxClient *NotifyInboxService
