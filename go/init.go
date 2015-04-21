package client

import (
	"fmt"
	"github.com/hprose/hprose-go/hprose"
	"gopkg.in/ini.v1"
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
	cfg, err = ini.Load(customPath)
	if err != nil {
		panic(err)
	}

	server := cfg.Section("").Key("server").MustString("http://127.0.0.1:3000/")

	fmt.Println("连接rpc服务器 :", server)

	// c := hprose.NewClient("tcp://127.0.0.1:3456/")
	c := hprose.NewClient(server)
	c.UseService(&UserClient)
	c.UseService(&StrategyStockClient)
	c.UseService(&StrategyUserClient)
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
