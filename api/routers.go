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
		m.GET("/message", GetMessage)
		m.POST("/message", SendMessage)
		m.PUT("/message")
		m.DELETE("/message")
	}
	r.Run()
}
