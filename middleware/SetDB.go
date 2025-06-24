package SetDB

import (
	"github.com/gin-gonic/gin"
	"order_go/sql"
)

// InitDB 保存数据库DB，用于全局访问
func InitDB() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", sql.GetDB())
		c.Next()
	}
}
