package main

import (
	"log"
	"runtime"
)

func runWithRecovery() {
	defer func() {
		if r := recover(); r != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			log.Printf("cron: panic running job: %v\n%s", r, buf)
		}
	}()
	// do something... panic
}

