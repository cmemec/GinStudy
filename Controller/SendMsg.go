package controller

import (
	model "GinStudy/Model"
	tool "GinStudy/Tool"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

type paramJson struct {
	Phone string `json:"phone"`
}

type SendMsg struct {
}

// 路由
func (sm *SendMsg) Router(engine *gin.Engine) {
	engine.Any("/sendMsg", sm.SendMsg)
}

// 发送短信
func (sm *SendMsg) SendMsg(context *gin.Context) {
	pj := new(paramJson)
	err := context.BindJSON(pj)
	if err != nil {
		log.Println(err.Error())
	}
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100000))

	var smInfo model.SmsCode
	bool, err := tool.XormEngine.Where("phone", pj.Phone).Get(&smInfo)
	if err != nil {
		tool.Failed(context, "")
	}
	log.Println(bool)
	addInfo := &model.SmsCode{
		Phone:      pj.Phone,
		Code:       code,
		CreateTime: time.Now().Unix(),
	}
	id, _ := tool.XormEngine.Insert(addInfo)
	if id > 0 {
		tool.Success(context, "")
	} else {
		tool.Failed(context, "")
	}

}
