package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"message-board/model"
	"message-board/service"
	"message-board/util"
)

func Register(c *gin.Context) {
	userName := c.PostForm("name")
	password := c.PostForm("password")
	if userName == "" || password == "" {
		util.ResParamError(c)
		return
	}
	u, err := service.SearchUserByName(userName)
	if err != nil && err != sql.ErrNoRows {
		util.RespInternalError(c)
		return
	}
	if u.UserName != "" {
		util.NormError(c, 300, "账户已存在")
		return
	}
	err = service.CreateUser(model.User{
		UserName: userName,
		Password: password,
	})
	if err != nil {
		util.RespInternalError(c)
		return
	}
	util.ResOk(c)
}

func Login(c *gin.Context) {
	userName := c.PostForm("name")
	password := c.PostForm("password")
	if userName == "" || password == "" {
		util.ResParamError(c)
		return
	}
	u, err := service.SearchUserByName(userName)
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
	if u.UserName != userName {
		util.NormError(c, 20001, "密码错误")
		return
	}
	c.SetCookie("name", userName, 1110, "/", "localhost", false, true)
	util.ResOk(c)
}
