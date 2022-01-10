package tool

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS int = 0
	FAILED  int = -1
)

// 成功
func Success(context *gin.Context, data interface{}) {
	context.JSON(http.StatusOK, gin.H{
		"code": SUCCESS,
		"msg":  "成功",
		"data": data,
	})
}

// 失败
func Failed(context *gin.Context, data interface{}) {
	context.JSON(http.StatusOK, gin.H{
		"code": FAILED,
		"msg":  "失败",
		"data": data,
	})
}
