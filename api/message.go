package api

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"message-board/middleware"
	"message-board/model"
	"message-board/service"
	"message-board/util"

	"github.com/gin-gonic/gin"
)

func GetMessage(c *gin.Context) {
	c.PostForm("mid")
}
func SendMessage(c *gin.Context) {
	userName, b := c.Get(middleware.CtxUser)
	if b != true {
		fmt.Println("未得到uid")
		return
	}
	recUid := c.PostForm("recUID")
	detail := c.PostForm("detail")
	if detail == "" || recUid == "" {
		util.ResParamError(c)
		return
	}
	_, err := service.SearchUserByName(recUid)
	if err != nil {
		if err == sql.ErrNoRows {
			util.NormError(c, 300, "用户不存在")
		} else {
			log.Printf("search user error:%v", err)
			util.RespInternalError(c)
			return
		}
		return
	}
	err = service.Comment(model.Message{
		SendName: userName.(string),
		RecUName: recUid,
		Detail:   detail,
		Time:     time.Now(),
	})
	if err != nil {
		log.Printf("error:%v", err)
		util.RespInternalError(c)
		return
	}
	util.ResOk(c)
}
func DeleteMessage(c *gin.Context) {
	userName, b := c.Get(middleware.CtxUser)
	if b != true {
		fmt.Println("未得到uid")
		return
	}
	err := service.Delete(userName.(string))
	if err != nil {
		log.Printf("error:%v", err)
		util.RespInternalError(c)
		return
	}
	util.ResOk(c)
}
