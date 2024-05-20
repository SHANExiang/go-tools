package jwt

import (
	"github.com/gin-gonic/gin"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Query("token")
		if token != "" {
			claims, err := ParseToken(token)
			if err == nil && time.Now().Unix() > claims.ExpiresAt{
				context.Abort()
			}
		}
		context.Next()
	}
}
