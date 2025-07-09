package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"order_go/routes"
	"order_go/sql"
)

func main() {
	router := gin.Default()

	sql.ConnectMysql()

	routes.InitRouter(router)

	// 404
	router.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "404 没有该接口",
		})
	})

	err := router.Run() // 监听并在 0.0.0.0:8080 上启动服务
	if err != nil {
		fmt.Println(err)
	}
}
