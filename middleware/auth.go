package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"message-board/util"
	"net/http"
)

const ctxUser string = "uid"

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		Username, err := c.Cookie("name")
		fmt.Println(Username)
		if err != nil {
			if err != http.ErrNoCookie {
				util.NormError(c, 300, "未登录")
				c.Abort()
				return
			} else {
				util.RespInternalError(c)
				c.Abort()
				return
			}
		}
		log.Println(Username)
		c.Set(ctxUser, Username)
		c.Next()
		return
	}
}
