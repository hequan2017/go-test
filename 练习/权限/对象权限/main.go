package main

import (
	"fmt"
	"github.com/casbin/casbin"
	"runtime"
)

type User struct {
	Id       int
	UserName string
	Group    []Group
}

type Group struct {
	Id       int
	Name     string
	App      string // app
	Type     string // 类型
	Method   string // 方法
	Priority int    // 优先级
}

type Obj struct {
	App    string // app
	Type   string // 类型
	Method string // 方法
}

func main() {
	osType := runtime.GOOS
	var path string
	if osType == "windows" {
		path = "权限\\对象权限\\conf\\abac_model.conf"
	} else if osType == "linux" {
		path = "权限/对象权限/conf/abac_model.conf"
	}

	e := casbin.NewEnforcer(path)

	group1 := Group{
		Name:     "group1",
		App:      "asset",
		Type:     "aliyun",
		Method:   "Get",
		Priority: 100,
	}

	group2 := Group{
		Name:     "group2",
		App:      "asset",
		Type:     "aliyun",
		Method:   "Get",
		Priority: 100,
	}

	//  用户 hequan  属于 group1 , group2
	user1 := User{
		UserName: "hequan",
		Group:    []Group{group1, group2},
	}

	obj := Obj{
		App:    "asset",
		Type:   "aliyun",
		Method: "Get",
	}

	var perms = false

	// 检查 用户 hequan 所有的组  是否有权限
	for _, v := range user1.Group {
		if e.Enforce(v, obj, "") {
			perms = true
			break
		}
	}
	if perms {
		fmt.Println("权限正常")
	} else {
		fmt.Println("没有权限")
	}
}
