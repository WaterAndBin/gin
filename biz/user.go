package biz

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	core "order_go/core"
	"order_go/data"
)

func UserLogin(context *gin.Context) {
	var login core.Login

	if err := context.ShouldBindJSON(&login); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"code": http.StatusInternalServerError, "error": "Missing or invalid parameters", "detail": "缺少必要的参数"})
		return
	}

	query := core.User{Account: login.Account}
	findResult := data.FindUser(context, query)

	if !findResult {
		//login.Password = utils.HashPassword(login.Password)
		// 没查找到数据
		data.InsertUser(context, login)
	} else {
		fmt.Println("666")
	}
}
