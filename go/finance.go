package client

var TFinanceShowKeys []string = []string{
	"fld_IncomePer",
	"fld_AssetsPer",
	"fld_AStockVol",
	"fld_StockVol",
	"fld_ProfitRate",
	"fld_FundPer",
	"fld_BStockVol",
	"fld_HStockVol",
	"fld_Assets",
	"fld_AssetsFlow",
	"fld_AssetsFixed",
	"fld_AssetsInvisible",
	"fld_DebtFlow",
	"fld_DebtLong",
	"fld_Fund",
	"fld_Rights",
	"fld_IncomeMain",
	"fld_ProfitMain",
	"fld_Profit",
	"fld_IncomeInvest",
	"fld_IncomeOther",
	"fld_ProfitSum",
	"fld_ProfitNet",
	"fld_ProfitRetain",
	"fld_ProfitPer",
	"fld_RightsRate",
	"fld_OperateCashNet",
	"fld_InvestCashNet",
	"fld_RaiseCashNet",
	"fld_IncomeMainIncRate",
	"fld_ProfitNetIncRate",
}

var TFinanceNames map[string]string = map[string]string{
	"fld_id":                "ID",
	"fld_Market":            "市场",
	"fld_Stock":             "股票代码",
	"fld_Date":              "日期",
	"fld_StockVol":          "总股本",
	"fld_AStockVol":         "流通A股",
	"fld_BStockVol":         "流通B股",
	"fld_HStockVol":         "流通H股",
	"fld_Assets":            "总资产",
	"fld_AssetsPer":         "每股净资",
	"fld_IncomePer":         "每股收益",
	"fld_AssetsFlow":        "流动资产",
	"fld_AssetsFixed":       "固定资产",
	"fld_AssetsInvisible":   "无形资产",
	"fld_ProfitRate":        "净资收益率",
	"fld_ProfitPer":         "每股未分配",
	"fld_DebtFlow":          "流动负债",
	"fld_DebtLong":          "长期负债",
	"fld_Rights":            "股东权益",
	"fld_RightsRate":        "股东权益比率",
	"fld_Fund":              "资本公积金",
	"fld_FundPer":           "每股公积金",
	"fld_OperateCashNet":    "经营现金流量",
	"fld_OperateCashIn":     "经营现金流入",
	"fld_OperateCashOut":    "经营现金流出",
	"fld_InvestCashNet":     "投资现金流量",
	"fld_InvestCashIn":      "投资现金流入",
	"fld_InvestCashOut":     "投资现金流出",
	"fld_RaiseCashNet":      "筹资现金流量",
	"fld_RaiseCashIn":       "筹资现金流入",
	"fld_RaiseCashOut":      "筹资现金流出",
	"fld_IncomeMain":        "主营收入",
	"fld_ProfitMain":        "主营利润",
	"fld_IncomeMainIncRate": "主营收入同比",
	"fld_Profit":            "营业利润",
	"fld_IncomeInvest":      "投资收益",
	"fld_IncomeOther":       "营业外收支",
	"fld_ProfitSum":         "利润总额",
	"fld_ProfitNet":         "净利润",
	"fld_ProfitRetain":      "未分配利润",
	"fld_ProfitNetIncRate":  "净利润同比",
}

type TFinance struct {
	Fld_id     string `xorm:"'fld_id'"`     // ID
	Fld_Market string `xorm:"'fld_Market'"` // 市场
	Fld_Stock  string `xorm:"'fld_Stock'"`  // 股票代码
	Fld_Date   string `xorm:"'fld_Date'"`   // 日期

	Fld_StockVol  string `xorm:"'fld_StockVol'"`  //总股本(万股)
	Fld_AStockVol string `xorm:"'fld_AStockVol'"` //流通A股(万股) 160
	Fld_BStockVol string `xorm:"'fld_BStockVol'"` //B股(万股)
	Fld_HStockVol string `xorm:"'fld_HStockVol'"` //H股(境外上市股)(万股)

	Fld_Assets          string `xorm:"'fld_Assets'"`          //总资产(千元)
	Fld_AssetsPer       string `xorm:"'fld_AssetsPer'"`       //每股净资(元)
	Fld_IncomePer       string `xorm:"'fld_IncomePer'"`       //每股收益(元)
	Fld_AssetsFlow      string `xorm:"'fld_AssetsFlow'"`      //流动资产(千元)
	Fld_AssetsFixed     string `xorm:"'fld_AssetsFixed'"`     //固定资产(千元)
	Fld_AssetsInvisible string `xorm:"'fld_AssetsInvisible'"` //无形资产(千元) 80
	Fld_ProfitRate      string `xorm:"'fld_ProfitRate'"`      //净资收益率(%)
	Fld_ProfitPer       string `xorm:"'fld_ProfitPer'"`       //每股未分配(元)

	Fld_DebtFlow string `xorm:"'fld_DebtFlow'"` //流动负债(千元)
	Fld_DebtLong string `xorm:"'fld_DebtLong'"` //长期负债(千元)

	Fld_Rights     string `xorm:"'fld_Rights'"`     //股东权益(千元)
	Fld_RightsRate string `xorm:"'fld_RightsRate'"` //股东权益比率(%)

	Fld_Fund    string `xorm:"'fld_Fund'"`    //资本公积金(千元)
	Fld_FundPer string `xorm:"'fld_FundPer'"` //每股公积金(元) 40

	Fld_OperateCashNet string `xorm:"'fld_OperateCashNet'"` //经营现金流量(千元)
	Fld_OperateCashOut string `xorm:"'fld_OperateCashOut'"` //经营现金流入(千元)
	Fld_OperateCashIn  string `xorm:"'fld_OperateCashIn'"`  //经营现金流出(千元)

	Fld_InvestCashNet string `xorm:"'fld_InvestCashNet'"` //投资现金流量(千元)
	Fld_InvestCashIn  string `xorm:"'fld_InvestCashIn'"`  //投资现金流入(千元)
	Fld_InvestCashOut string `xorm:"'fld_InvestCashOut'"` //投资现金流出(千元)

	Fld_RaiseCashNet string `xorm:"'fld_RaiseCashNet'"` //筹资现金流量(千元)
	Fld_RaiseCashIn  string `xorm:"'fld_RaiseCashIn'"`  //筹资现金流入(千元)
	Fld_RaiseCashOut string `xorm:"'fld_RaiseCashOut'"` //筹资现金流出(千元)

	Fld_IncomeMain        string `xorm:"'fld_IncomeMain'"`        //主营收入(千元) 120
	Fld_ProfitMain        string `xorm:"'fld_ProfitMain'"`        //主营利润(千元)
	Fld_IncomeMainIncRate string `xorm:"'fld_IncomeMainIncRate'"` //主营收入同比(%)

	Fld_Profit       string `xorm:"'fld_Profit'"`       //营业利润(千元)
	Fld_IncomeInvest string `xorm:"'fld_IncomeInvest'"` //投资收益(千元)
	Fld_IncomeOther  string `xorm:"'fld_IncomeOther'"`  //营业外收支(千元)

	Fld_ProfitSum    string `xorm:"'fld_ProfitSum'"`    //利润总额(千元)
	Fld_ProfitNet    string `xorm:"'fld_ProfitNet'"`    //净利润(千元)
	Fld_ProfitRetain string `xorm:"'fld_ProfitRetain'"` //未分配利润(千元)

	Fld_ProfitNetIncRate string `xorm:"'fld_ProfitNetIncRate'"` //净利润同比(%)

	Fld_Locked     string `xorm:"'fld_Locked'"`
	Fld_Deleted    string `xorm:"'fld_Deleted'"`
	Fld_UpdateTime string `xorm:"'fld_UpdateTime'"`
	Fld_TimeStamp  string `xorm:"'fld_TimeStamp'"`
}

type TFinanceService struct {
	FinanceInfoByCodeCache func(code string) (info map[string]string, err error)
	FinanceInfoByCode      func(code string) (info map[string]string, err error)
}

var TFinanceClient *TFinanceService
