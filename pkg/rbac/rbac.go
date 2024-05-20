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
	modelPath := currentDir + "\\model_rbac.conf"
	policyPath := currentDir + "\\policy_rbac.csv"
	enforcer, err := casbin.NewEnforcer(modelPath, policyPath)
	if err != nil {
		log.Fatalf("NewEnforcer failed:%v\n", err)
	}

	check(enforcer, "lizi", "data", "read")
	check(enforcer, "lizi", "data", "write")
	check(enforcer, "dajun", "data", "write")
	check(enforcer, "dajun", "data", "read")
	//lizi can read data
	//lizi can not write data
	//dajun can write data
	//dajun can read data

}
