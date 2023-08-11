package util

import (
	"github.com/gin-gonic/gin"
	"hasdsd.cn/op-dl-server/define"
	"log"
	"strconv"
)

// GetPageInfo 获取分页参数
func GetPageInfo(ctx *gin.Context) (int, int) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		log.Println("get param error", err.Error())
	}
	size, err := strconv.Atoi(ctx.DefaultQuery("size", define.DefaultSize))
	if err != nil {
		log.Println("get param error", err.Error())
	}
	page = (page - 1) * size
	return page, size
}
