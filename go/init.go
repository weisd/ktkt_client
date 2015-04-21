package client

import (
	"github.com/hprose/hprose-go/hprose"
)

type Wherer struct {
	Where string
	Args  []interface{}
}

type Limiter struct {
	Start  int
	Offset int
}

type Sorter struct {
	OrderBy string
}

func init() {
	// c := hprose.NewClient("tcp://127.0.0.1:3456/")
	c := hprose.NewClient("http://127.0.0.1:3000/")
	c.UseService(&UserClient)
	c.UseService(&StrategyStockClient)
	c.UseService(&StrategyUserClient)
}
