package quan

import (
	"time"
)

type KtktQuanThread struct {
	Id      int64  `json:"id"`
	TId     int64  `json:"t_id"`
	PId     int64  `json:"p_id"`
	CatesId int64  `json:"cates_id"`
	Title   string `json:"title"`
	Content string `json:"content"`

	Author       string    `json:"author"`
	Status       int       `json:"status"`
	DelStatus    int       `json:"del_status"`
	View         int64     `json:"view"`
	Comment      int64     `json:"comment"`
	Created      time.Time `json:"created"`
	Updated      time.Time `json:"udpated"`
	UpdatedUser  int64     `json:"updated_user"`
	Priority     int64     `json:"priority"`
	ThreadStatus int       `json:"thread_status"`
	ThreadTop    bool      `json:"thread_top"`
	Img          string    `json:"img"`
	Tags         string    `json:"tags"`

	CreatedAt time.Time `xorm:"created" json:"created_at"`
	UpdatedAt time.Time `xorm:"updated" json:"updated_at"`
	DeletedAt time.Time `xorm:"deleted" json:"deleted_at"`
}

type KtktQuanThreadService struct {
	QuanGetThreadInfoById  func(id int64) (*KtktQuanThread, error)
	QuanGetThreadInfoByTid func(id string) (*KtktQuanThread, error)
}

var KtktQuanThreadClient *KtktQuanThreadService
