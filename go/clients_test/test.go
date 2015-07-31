package main

import (
	client "../../go"
	"errors"
	"fmt"
)

func main() {
	client.InitClient("http://127.0.0.1:3030")

	// f := client.Wherer{}
	// f.Where = "id = ?"
	// f.Args = []interface{}{90}
	// res, err := client.StrategyClient.StrategyGetOneCache(f)
	// if err != nil {
	// 	panic(err)
	// }

	// TestNotification()

	// fmt.Println(res)
	// 创建
	// TestKtRole()
	// TestKtPermission()
	// TestRbac()
	// TestUser()

	// TestStockSend()
	TestFutures()
}

func TestFutures() {
	res, err := client.FuturesClient.FuturesInfoByCode("IF1508")
	fmt.Println(res, err)

	res, err = client.FuturesClient.FuturesInfoByCodeNFields("IF1508", []string{"Name"})
	fmt.Println(res, err)

}

func TestStockSend() {
	// []string{"10000001530", "10000600430", "10000600730", "10000660830", "10000600130", "10000661130", "10011930330"}
	sidAll, err := client.RbacClient.GetAllStrategyNodeIds()

	fmt.Println(sidAll, err)
	aft, err := client.RbacClient.DelStockSendUserByStrategyId("10000001530", 123456)
	fmt.Println(aft, err)

	for _, k := range sidAll {
		// 开启发送功能的用户ids
		uids := client.RbacClient.GetStockSendUidsByStrategyId(k)
		fmt.Println(k, uids)
	}

	// // 删除发送功能

	// // 删除发送功能
	client.RbacClient.AddStockSendUserByStrategyId("10000001530", []int64{1113152238, 123456})
	// // 关闭用户的所有
	client.RbacClient.ClearStockSendByUid(123456)
	// // 所有策略
	for _, k := range sidAll {
		// 开启发送功能的用户ids
		uids := client.RbacClient.GetStockSendUidsByStrategyId(k)
		fmt.Println(k, uids)
	}

}

func TestNotification() {
	fmt.Println("TestNotification")

	// n := new(client.Notification)
	// n.CategoryId = client.MOTIFICATION_LIVE
	// n.Content = "测试消息"
	// n.Resource = "test"
	// n.ResId = 1
	// // m := client.Notification(n)
	// // if {
	// // 	fmt.Println("err")
	// // 	return
	// // }

	// aft, err := client.NotificationClient.NotificationAdd(n)
	// fmt.Println(aft, err)

	// fmt.Println(n)
	// // return

	// insertId, err := client.NotificationClient.NotificationCreate(n)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println("insertId", insertId)

	// res, err := client.NotificationClient.NotificationGetOne(client.Wherer{Where: "id = ?", Args: []interface{}{insertId}})
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println("res", res)

	// list, err := client.NotificationClient.NotificationGetList(client.Wherer{}, client.Limiter{}, client.Sorter{})
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println("list", list)
}

func TestUser() {
	// user, err := client.UserClient.UserGetInfoByPhone("")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(user)

	// user, err := client.UserClient.UserGetInfoByName("dada")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(user)

	// ok, err := client.UserClient.UserCheckPassword(user.Id, "")
	// if err != nil {
	// 	panic(err)
	// }

	// //"ICYPH2I5s2BlKFUij+JGBaI7EyVwcklV"

	// fmt.Println("ok ? ", ok)

	// user, err = client.UserClient.UserGetInfoByEmail("")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(user)
}

func TestRbac() {

	// 用户的所有策略权限
	// list, err := client.RbacClient.GetUserNodeIds(3, 1)
	// if err != nil {
	// 	panic(err)
	// }

	// for _, v := range list {
	// 	fmt.Println(v)
	// }
	//

	// ok, err := client.RbacClient.HasRoles(1, "feng")
	// ok, err := client.RbacClient.HasNode(1, 2001, 1)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(ok)
}

func TestKtPermission() {

}

func TestKtRole() {

	n := &client.KtRoles{}
	n.Title = "测一"
	n.Keywords = "test1"

	id, err := client.KtRolesClient.KtRolesCreate(n)
	if err != nil {
		panic(err)
	}

	fmt.Printf("插入数据ok , id : %d \n", id)

	// 列表
	where := client.Wherer{}
	limit := client.Limiter{}
	sort := client.Sorter{}

	list, err := client.KtRolesClient.KtRolesGetList(where, limit, sort)
	if err != nil {
		panic(err)
	}

	fmt.Println("列表：", list)

	// 查询
	where = client.Wherer{}
	where.Where = "id=?"
	where.Args = []interface{}{id}

	info, err := client.KtRolesClient.KtRolesGetOne(where)
	if err != nil {
		panic(err)
	}

	if info == nil || info.Keywords != "test1" {
		panic(errors.New("查询数据不对"))
	}
	// 修改
	update := map[string]interface{}{}
	update["title"] = "测试修改"

	aft, err := client.KtRolesClient.KtRolesUpdate(id, update)
	if err != nil {
		panic(err)
	} else if aft == 0 {
		panic(errors.New("修改失败"))
	}

	where = client.Wherer{}
	where.Where = "title=?"
	where.Args = []interface{}{"测试修改"}

	info, err = client.KtRolesClient.KtRolesGetOne(where)
	if err != nil {
		panic(err)
	}

	if info == nil || info.Title != "测试修改" {
		panic(errors.New("查询修改数据不对"))
	}

	// 删除
	where = client.Wherer{}
	where.Where = "title=?"
	where.Args = []interface{}{"测试修改"}
	aft, err = client.KtRolesClient.KtRolesDelete(where)
	if err != nil {
		panic(err)
	} else if aft == 0 {
		panic(errors.New("删除失败"))
	}

	where = client.Wherer{}
	where.Where = "id=?"
	where.Args = []interface{}{id}

	info, err = client.KtRolesClient.KtRolesGetOne(where)
	if err != nil {
		panic(err)
	}

	if info != nil {
		panic(errors.New("删除失败"))
	}

	fmt.Println("test ok")
}
