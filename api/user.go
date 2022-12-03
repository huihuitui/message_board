package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"message-board/middleware"
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
		fmt.Printf("errererer:%v", err)
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
		log.Printf("Encrypt failed:%v", err)
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
	enRes, err := util.Encrypt(password)
	if err != nil {
		log.Printf("Encrypt failed:%v", err)
		util.RespInternalError(c)
		return
	}
	if u.Password != enRes {
		util.NormError(c, 20001, "密码错误")
		return
	}
	token, err := middleware.GenRegisteredClaims(userName)
	if err != nil {
		log.Printf("get jwt failed:%v", err)
		util.NormError(c, 200, "获取token失败")
		return
	}
	util.RspToken(c, 200, token)
}
