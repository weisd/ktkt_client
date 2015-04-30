package client

type MystockType struct {
	Id    int64
	Uid   int64
	Title string
}

type MystockTypeService struct {

	// 取列表 带缓存
	MystockTypeGetListCache func(where Wherer, limit Limiter, sort Sorter) ([]MystockType, error)

	// 取列表
	MystockTypeGetList func(where Wherer, limit Limiter, sort Sorter) ([]MystockType, error)

	// 取信息
	MystockTypeGetOneCache func(where Wherer) (*MystockType, error)

	// 取信息
	MystockTypeGetOne func(where Wherer) (*MystockType, error)

	// 创建信息
	MystockTypeCreate func(kt_roles *MystockType) (int64, error)

	// 删除 单条
	MystockTypeDelete func(where Wherer) (int64, error)

	// 更新数据
	MystockTypeUpdate func(id int64, updater map[string]interface{}) (int64, error)

	// 消除缓存
	MystockTypeFlushCache func() error
}

var MystockTypeClient *MystockTypeService
