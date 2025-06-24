package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	SetDB "order_go/middleware"
	"order_go/routes"
	"order_go/sql"
)

func main() {
	router := gin.Default()

	sql.ConnectMysql()
	router.Use(SetDB.InitDB())

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
