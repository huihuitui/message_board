package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type respTemplate struct {
	Status int    `json:"status"`
	Info   string `json:"info"`
}

var OK = respTemplate{
	Status: 200,
	Info:   "suceess",
}

var ParamError = respTemplate{
	Status: 300,
	Info:   "param error",
}

var InternalError = respTemplate{
	Status: 500,
	Info:   "internal error",
}

func ResOk(c *gin.Context) {
	c.JSON(http.StatusOK, OK)
}

func ResParamError(c *gin.Context) {
	c.JSON(http.StatusBadRequest, ParamError)
}

func RespInternalError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, InternalError)
}
func NormError(c *gin.Context, status int, info string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status": status,
		"Info":   info,
	})
}
func RspToken(c *gin.Context, status int, info string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status": status,
		"token":  info,
	})
}
