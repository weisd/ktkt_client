package client

import (
	"fmt"
	"github.com/hprose/hprose-go/hprose"
	// "gopkg.in/ini.v1"
	// "os"
	// "os/exec"
	// "path"
	// "path/filepath"
	"reflect"
	// "strings"
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

func InitClient(server string) {
	// workDir, err := WorkDir()
	// if err != nil {
	// 	panic(err)
	// }
	// customPath := path.Join(workDir, "custom/conf/rpc.ini")
	// cfg, err := ini.Load(customPath)
	// if err != nil {
	// 	panic(err)
	// }

	// server := cfg.Section("").Key("server").MustString("http://127.0.0.1:3000/")

	fmt.Println("连接rpc服务器 :", server)

	hprose.ClassManager.Register(reflect.TypeOf(User{}), "User", "json")
	hprose.ClassManager.Register(reflect.TypeOf(Strategy{}), "Strategy", "json")
	hprose.ClassManager.Register(reflect.TypeOf(StrategyStock{}), "StrategyStock", "json")

	hprose.ClassManager.Register(reflect.TypeOf(KtRoles{}), "KtRoles", "json")
	hprose.ClassManager.Register(reflect.TypeOf(KtPermissions{}), "KtPermissions", "json")

	hprose.ClassManager.Register(reflect.TypeOf(Mystock{}), "Mystock", "json")
	hprose.ClassManager.Register(reflect.TypeOf(MystockGroup{}), "MystockGroup", "json")

	hprose.ClassManager.Register(reflect.TypeOf(Notification{}), "Notification", "json")
	hprose.ClassManager.Register(reflect.TypeOf(NotifyInbox{}), "NotifyInbox", "json")

	hprose.ClassManager.Register(reflect.TypeOf(Clients{}), "Clients", "json")
	hprose.ClassManager.Register(reflect.TypeOf(Stock{}), "Stock", "json")
	hprose.ClassManager.Register(reflect.TypeOf(StrategyInfo{}), "StrategyInfo", "json")
	hprose.ClassManager.Register(reflect.TypeOf(TFinance{}), "TFinance", "json")

	// c := hprose.NewClient("tcp://127.0.0.1:3456/")
	c := hprose.NewClient(server)
	c.UseService(&UserClient)
	c.UseService(&StrategyClient)
	c.UseService(&StrategyStockClient)
	c.UseService(&StrategyUserClient)

	c.UseService(&KtRolesClient)
	c.UseService(&KtPermissionsClient)
	c.UseService(&KtRoleUserClient)
	c.UseService(&KtRolePermissionClient)

	c.UseService(&RbacClient)

	c.UseService(&MystockClient)
	c.UseService(&MystockGroupClient)

	c.UseService(&StockClient)

	c.UseService(&ClientStubClient)

	c.UseService(&NotificationClient)
	c.UseService(&NotifyInboxClient)

	c.UseService(&ClientsClient)

	c.UseService(&StrategyInfoClient)

	c.UseService(&FuturesClient)
	c.UseService(&TFinanceClient)

	// c.AddFilter(LogFilter{})
}

// func WorkDir() (string, error) {
// 	execPath, err := ExecPath()
// 	return path.Dir(strings.Replace(execPath, "\\", "/", -1)), err
// }

// func ExecPath() (string, error) {
// 	file, err := exec.LookPath(os.Args[0])
// 	if err != nil {
// 		return "", err
// 	}
// 	p, err := filepath.Abs(file)
// 	if err != nil {
// 		return "", err
// 	}
// 	return p, nil
// }
