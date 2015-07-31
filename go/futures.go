package client

type Futures struct {
	InstrumentID    string
	Market          string
	Security        string
	Name            string
	PriceTick       string
	VolumeMultiple  string
	DelivDate       string
	TradingDay      string
	PreSettle       string
	PreClose        string
	PreOpenInterest string
	UpperLimit      string
	LowerLimit      string

	Time         string
	Open         string
	High         string
	Low          string
	Last         string
	Volume       string
	Turnover     string
	OpenInterest string
	Settle       string
	Bid          string
	Ask          string
	BidVolume    string
	AskVolume    string
}

var FuturesKeys = []string{
	"InstrumentID",
	"Market",
	"Security",
	"Name",
	"PriceTick",
	"VolumeMultiple",
	"DelivDate",
	"TradingDay",
	"PreSettle",
	"PreClose",
	"PreOpenInterest",
	"UpperLimit",
	"LowerLimit",
	"Time",
	"Open",
	"High",
	"Low",
	"Last",
	"Volume",
	"Turnover",
	"OpenInterest",
	"Settle",
	"Bid",
	"Ask",
	"BidVolume",
	"AskVolume",
}

type FuturesService struct {
	FuturesInfoByCode        func(id string) (info map[string]string, err error)
	FuturesInfoByCodeNFields func(id string, fields []string) (info map[string]string, err error)
}

var FuturesClient *FuturesService
