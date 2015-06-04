package quan

import (
	"time"
)

type KtktQuanComment struct {
	Id        int64  `json:"id"`
	TId       string `json:"t_id"`
	UserId    int64  `json:"user_id"`
	CommentId string `json:"comment_id"`
	Support   int64  `json:"support"`
	TStatus   bool   `json:"t_status"`
	Pid       string `json:"pid"`
	Content   string `json:"content"`

	Created int64 `json:"created"`

	CreatedAt time.Time `xorm:"created" json:"created_at"`
	UpdatedAt time.Time `xorm:"updated" json:"updated_at"`
	DeletedAt time.Time `xorm:"deleted" json:"deleted_at"`
}

type KtktQuanCommentService struct {
	QuanGetCommentInfoById func(id int64) (*KtktQuanComment, error)
}

var KtktQuanCommentClient *KtktQuanCommentService
