package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(context *gin.Context, data interface{}) {
	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": data,
	})
}

func Error(context *gin.Context, msg string) {
	context.JSON(http.StatusInternalServerError, gin.H{
		"code": http.StatusInternalServerError,
		"msg":  msg,
	})
}
