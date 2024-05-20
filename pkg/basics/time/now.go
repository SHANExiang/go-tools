package main

import (
	"fmt"
	"time"
)

func main() {
    location := time.Now().Location()
    fmt.Println(location)     //Local
    now := time.Now().In(location)
    fmt.Println(now)           // now
    fmt.Println(now.Format("2006-01-02_15:04:05"))
}
