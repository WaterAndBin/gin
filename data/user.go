package data

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"order_go/core"
	"order_go/utils"
	"time"
)

func InsertUser(context *gin.Context, login core.Login) {
	db := context.MustGet("db").(*gorm.DB)
	createResult := db.Create(&core.User{
		Account:  login.Account,
		Password: login.Password,
		UserName: time.Now().String(),
	})

	if createResult.Error != nil {
		utils.Error(context, "user写入错误")
	} else {
		utils.Success(context, "user写入成功")
	}
}

func FindUser(context *gin.Context, query core.User) bool {
	db := context.MustGet("db").(*gorm.DB)
	var user core.User

	result := db.Where(query).Limit(1).First(&user)

	if result.Error != nil {
		fmt.Println(result.Error)
		return false
	}
	// 没查找到数据
	if result.RowsAffected == 0 {
		return false
	}
	return true
}
