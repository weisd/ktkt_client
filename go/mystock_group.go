package client

type MystockGroup struct {
	Id    int64  `json:"id"`
	Uid   int64  `json:"uid"`
	Title string `json:"title"`
}

type MystockGroupService struct {

	// 取列表 带缓存
	MystockGroupGetListCache func(where Wherer, limit Limiter, sort Sorter) ([]MystockGroup, error)

	// 取列表
	MystockGroupGetList func(where Wherer, limit Limiter, sort Sorter) ([]MystockGroup, error)

	// 取信息
	MystockGroupGetOneCache func(where Wherer) (*MystockGroup, error)

	// 取信息
	MystockGroupGetOne func(where Wherer) (*MystockGroup, error)

	// 创建信息
	MystockGroupCreate func(kt_roles *MystockGroup) (int64, error)

	// 删除 单条
	MystockGroupDelete func(where Wherer) (int64, error)

	// 更新数据
	MystockGroupUpdate func(id int64, updater map[string]interface{}) (int64, error)

	// 消除缓存
	MystockGroupFlushCache func() error
}

var MystockGroupClient *MystockGroupService
