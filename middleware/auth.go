package middleware

import (
	"github.com/gin-gonic/gin"
	"message-board/util"
	"net/http"
)

const CtxUser string = "uid"

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		Username, err := c.Cookie("name")
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
		c.Set(CtxUser, Username)
		c.Next()
		return
	}
}
