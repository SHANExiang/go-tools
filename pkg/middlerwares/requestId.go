package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goprojects/pkg/uuidUtil"
)

func Setup() gin.HandlerFunc {
	return func(context *gin.Context) {
		requestId := context.Request.Header.Get("X-Request-Id")
		if requestId == "" {
			requestId = uuidUtil.GenerateSnowflake()
		}
		context.Set("X-Request-Id", requestId)
		context.Writer.Header().Set("X-Request-Id", requestId)
		context.Next()
	}
}

func main() {
	fmt.Println(uuidUtil.GenerateSnowflake())
}