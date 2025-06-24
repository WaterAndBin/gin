package routes

import (
	"github.com/gin-gonic/gin"
	"order_go/biz"
)

// 初始化路由
func InitRouter(router *gin.Engine) {
	{
		v1 := router.Group("/user")
		v1.POST("/login", biz.UserLogin)
	}
}
