package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	uuidUtil2 "go-tools/utils/uuidUtil"
)

func Setup() gin.HandlerFunc {
	return func(context *gin.Context) {
		requestId := context.Request.Header.Get("X-Request-Id")
		if requestId == "" {
			requestId = uuidUtil2.GenerateSnowflake()
		}
		context.Set("X-Request-Id", requestId)
		context.Writer.Header().Set("X-Request-Id", requestId)
		context.Next()
	}
}

func main() {
	fmt.Println(uuidUtil2.GenerateSnowflake())
}