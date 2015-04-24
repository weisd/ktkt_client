package client

import (
	"fmt"
	"github.com/hprose/hprose-go/hprose"
	"gopkg.in/ini.v1"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"reflect"
	"strings"
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
	workDir, err := WorkDir()
	if err != nil {
		panic(err)
	}
	customPath := path.Join(workDir, "custom/conf/rpc.ini")
	cfg, err := ini.Load(customPath)
	if err != nil {
		panic(err)
	}

	server := cfg.Section("").Key("server").MustString("http://127.0.0.1:3000/")

	fmt.Println("连接rpc服务器 :", server)

	hprose.ClassManager.Register(reflect.TypeOf(User{}), "User", "json")
	hprose.ClassManager.Register(reflect.TypeOf(Strategy{}), "Strategy", "json")
	hprose.ClassManager.Register(reflect.TypeOf(StrategyStock{}), "StrategyStock", "json")

	hprose.ClassManager.Register(reflect.TypeOf(KtRoles{}), "KtRoles", "json")
	hprose.ClassManager.Register(reflect.TypeOf(KtPermissions{}), "KtPermissions", "json")

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
}

func WorkDir() (string, error) {
	execPath, err := ExecPath()
	return path.Dir(strings.Replace(execPath, "\\", "/", -1)), err
}

func ExecPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	p, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	return p, nil
}
