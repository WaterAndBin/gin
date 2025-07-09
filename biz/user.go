package biz

import (
	"github.com/gin-gonic/gin"
	"net/http"
	core "order_go/core"
	"order_go/data"
	"order_go/utils"
)

func UserLogin(context *gin.Context) {
	var login core.Login

	if err := context.ShouldBindJSON(&login); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"code": http.StatusInternalServerError, "error": "Missing or invalid parameters", "detail": "缺少必要的参数"})
		return
	}
	if login.Password == "" {
		utils.Error(context, "密码不能为空")
		return
	}

	user, findResult := data.FindUser(context, core.User{Account: login.Account})

	if !findResult {
		//login.Password = utils.HashPassword(login.Password)
		// 没查找到数据，就插入一条新的数据
		data.InsertUser(context, login)
	} else {
		data.GetUserInfo(context, login, user)
	}
}
