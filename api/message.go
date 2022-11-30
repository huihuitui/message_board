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
	_, b := c.Get(middleware.CtxUser)
	if b != true {
		fmt.Println("未得到uid")
		return
	}
	m, err := service.Get()
	if err != nil {
		log.Printf("error:%v", err)
		util.NormError(c, 20002, "获取失败")
		return
	}
	for _, v := range m {
		c.JSON(200, v)
	}
}
func GetComment(c *gin.Context) {
	_, b := c.Get(middleware.CtxUser)
	if b != true {
		fmt.Println("未得到uid")
		return
	}
	userName := c.Query("re_uid")
	m, err := service.GetComment(userName)
	if err != nil {
		log.Printf("error:%v", err)
		util.NormError(c, 20002, "获取失败")
		return
	}
	for _, v := range m {
		c.JSON(200, v)
	}
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
	m, err := service.SearchUserByName(recUid)
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
		UID:      m.ID,
		SendName: userName.(string),
		RecUName: recUid,
		Detail:   detail,
		Time:     time.Now(),
	})
	if err != nil {
		log.Printf("error:%v", err)
		util.NormError(c, 20004, "评论失败")
		return
	}
	util.ResOk(c)
}
func AlterMessage(c *gin.Context) {
	userName, b := c.Get(middleware.CtxUser)
	if b != true {
		fmt.Println("未得到uid")
		return
	}
	m, err := service.GetUserMessage(userName.(string))
	if err != nil {
		log.Printf("error:%v", err)
		util.NormError(c, 20004, "获取失败")
		return
	}
	for _, v := range m {
		c.JSON(200, v)
	}
}
func ModifyMessage(c *gin.Context) {
	userName, b := c.Get(middleware.CtxUser)
	if b != true {
		fmt.Println("未得到uid")
		return
	}
	newDetail := c.PostForm("newDetail")
	detail := c.PostForm("detail")
	err := service.ModifyMessage(userName.(string), newDetail, detail)
	if err != nil {
		log.Printf("error:%v", err)
		util.NormError(c, 20003, "修改失败")
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
		util.NormError(c, 20002, "删除失败")
		return
	}
	util.ResOk(c)
}
