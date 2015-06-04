package quan

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

	// server := cfg.Section("").Key("server").MustString("http://127.0.0.1:3000/")

	fmt.Println("连接rpc服务器 :", server)

	hprose.ClassManager.Register(reflect.TypeOf(Orders{}), "Orders", "json")
	hprose.ClassManager.Register(reflect.TypeOf(KtktQuanComment{}), "KtktQuanComment", "json")
	hprose.ClassManager.Register(reflect.TypeOf(KtktQuanThread{}), "KtktQuanThread", "json")

	// c := hprose.NewClient("tcp://127.0.0.1:3456/")
	c := hprose.NewClient(server)
	c.UseService(&OrdersClient)
	c.UseService(&KtktQuanCommentClient)
	c.UseService(&KtktQuanThreadClient)

	// c.AddFilter(LogFilter{})
}
