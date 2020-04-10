package Weblib

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-micro-pro/HTTP-API/Services"
)

func InitMiddleware(prodService Services.ProdService) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Keys = make(map[string]interface{})
		context.Keys["prodservrice"] = prodService //赋值
		context.Next()
	}
}

//异常处理
func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				c.JSON(500, gin.H{
					"status": fmt.Sprintf("%s", r),
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
