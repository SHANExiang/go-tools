package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"log"
	"path"
	"runtime"
)

func GetRootPath() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

func check(enforcer *casbin.Enforcer, sub, obj, act string)  {
	ok, _ := enforcer.Enforce(sub, obj, act)
	if ok {
		fmt.Printf("%s can %s %s\n", sub, act, obj)
	} else {
		fmt.Printf("%s can not %s %s\n", sub, act, obj)
	}
}

func main() {
	currentDir := GetRootPath()
	modelPath := currentDir + "\\model.conf"
	policyPath := currentDir + "\\policy.csv"
    enforcer, err := casbin.NewEnforcer(modelPath, policyPath)
    if err != nil {
    	log.Fatalf("NewEnforcer failed:%v\n", err)
	}

	check(enforcer, "dajun", "data1", "read")
	check(enforcer, "lizi", "data2", "write")
	check(enforcer, "dajun", "data1", "write")
	check(enforcer, "dajun", "data2", "read")
    //dajun can read data1
	//lizi can write data2
	//dajun can not write data1
	//dajun can not read data2

	check(enforcer, "root", "data1", "read")
	check(enforcer, "root", "data2", "write")
	check(enforcer, "root", "data1", "write")
	check(enforcer, "root", "data2", "read")
    //root can read data1
	//root can write data2
	//root can write data1
	//root can read data2
	// 超级管理员可以进行任何操作

}
