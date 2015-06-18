package client

// 股票信息
type Stock struct {
	HQZQDM  string
	HQZQJC  string
	HQJRKP  string
	HQZRSP  string
	HQZJCJ  string
	HQZGCJ  string
	HQZDCJ  string
	HQBJW1  string
	HQSJW1  string
	HQCJSL  string
	HQCJJE  string
	HQBSL1  string
	HQBSL2  string
	HQBJW2  string
	HQBSL3  string
	HQBJW3  string
	HQBSL4  string
	HQBJW4  string
	HQBSL5  string
	HQBJW5  string
	HQSSL1  string
	HQSSL2  string
	HQSJW2  string
	HQSSL3  string
	HQSJW3  string
	HQSSL4  string
	HQSJW4  string
	HQSSL5  string
	HQSJW5  string
	HQSYL1  string
	Code    string
	Date    string
	Time    string
	Delete  string
	AddTime string
}

var CodeKeys = []string{
	"HQZQDM",
	"HQZQJC",
	"HQJRKP",
	"HQZRSP",
	"HQZJCJ",
	"HQZGCJ",
	"HQZDCJ",
	"HQBJW1",
	"HQSJW1",
	"HQCJSL",
	"HQCJJE",
	"HQBSL1",
	"HQBSL2",
	"HQBJW2",
	"HQBSL3",
	"HQBJW3",
	"HQBSL4",
	"HQBJW4",
	"HQBSL5",
	"HQBJW5",
	"HQSSL1",
	"HQSSL2",
	"HQSJW2",
	"HQSSL3",
	"HQSJW3",
	"HQSSL4",
	"HQSJW4",
	"HQSSL5",
	"HQSJW5",
	"HQSYL1",
	"code",
	"date",
	"time",
	"delete",
	"add_time",
}

type StockService struct {
	GetInfoByCodeNfields func(code string, fields []string) (map[string]string, error)
	// 取股票信息
	GetInfoByCode func(code string) (*Stock, error)
	// 取股票名称
	GetName func(code string) (string, error)

	// 取股票字段值
	GetField func(code, field string) (string, error)
}

var StockClient *StockService
