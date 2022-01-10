package controller

import (
	"github.com/gin-gonic/gin"
)

type User struct {
}

// 路由
func (user *User) Router(engine *gin.Engine) {
	engine.GET("/login", user.login)
}

// 方法
func (user *User) login(context *gin.Context) {

}
