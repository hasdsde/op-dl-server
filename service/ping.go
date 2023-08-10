package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Ping
// @Summary ping
// @Description ping
// @Tags 测试
// @Success 200 {string} json "{"code":"200","data":"pong"}"
// @Router /ping [get]
func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "pong",
	})
}
