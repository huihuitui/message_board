package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"message-board/model"
	"message-board/service"
	"message-board/util"
)

func GetMessage(c *gin.Context) {
	c.PostForm("mid")
}
func SendMessage(c *gin.Context) {
	userName, b := c.Get("uid")
	if b != true {
		fmt.Println("未得到uid")
		return
	}
	recUid := c.PostForm("recUID")
	detail := c.PostForm("detail")
	_, err := service.SearchUserByName(recUid) //查看回复用户是否存在
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
	v, _ := userName.(string)
	err = service.Comment(model.Message{
		SendName: userName.(string),
		RecUName: recUid,
		Detail:   detail,
	})
	fmt.Println(v)
	if err != nil {
		log.Printf("error:%v", err)
		util.RespInternalError(c)
		return
	}
	util.ResOk(c)
}
