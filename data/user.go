package data

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"order_go/core"
	"order_go/sql"
	"order_go/utils"
	"time"
)

func InsertUser(context *gin.Context, login core.Login) {
	hashRes, err := utils.HashPassword(login.Password)

	if err != nil {
		utils.Error(context, "出错！")
	}

	createResult := sql.DB.Create(&core.User{
		Account:  login.Account,
		Password: hashRes,
		UserName: time.Now().String(),
	})

	if createResult.Error != nil {
		utils.Error(context, "user写入错误")
	} else {
		utils.Success(context, "user写入成功")
	}
}

func FindUser(context *gin.Context, query core.User) (core.User, bool) {
	var user core.User

	result := sql.DB.Where(query).Limit(1).First(&user)

	if result.Error != nil {
		fmt.Println(result.Error)
		return user, false
	}
	// 没查找到数据
	if result.RowsAffected == 0 {
		return user, false
	}
	return user, true
}

func GetUserInfo(context *gin.Context, login core.Login, query core.User) {
	err := utils.ComparePasswords(query.Password, login.Password)
	if err != nil {
		utils.Error(context, "账号或者密码有误")
		return
	}
	utils.Success(context, "登录成功")
}
