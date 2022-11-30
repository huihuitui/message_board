package api

import (
	"github.com/gin-gonic/gin"
	"message-board/middleware"
)

func InitRouter() {
	r := gin.Default()
	u := r.Group("/user")
	{
		u.POST("/register", Register)
		u.POST("/login", Login)
		u.POST("/password")
	}
	m := r.Group("/message")
	{
		m.Use(middleware.Auth())
		m.GET("/message", GetMessage)        //查看所有留言板
		m.GET("/comment", GetComment)        //查看一个留言及其评论
		m.GET("/altermessage", AlterMessage) //用户先查看自己所有评论，根据评论内容修改某一条评论
		m.POST("/message", SendMessage)
		m.PUT("/message", ModifyMessage)
		m.DELETE("/message", DeleteMessage)
	}
	r.Run()
}
