package main

import (
	"fmt"
	"regexp"
)

func main() {
	pattern := "a[a-z]b*"
	s := "acddb"
	compile, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Printf("Fail to compile pattern:%v", err)
	}
	res := compile.FindAllString(s, 4)
	fmt.Println(res)
}
