package util

import (
	"github.com/gin-gonic/gin"
	"hasdsd.cn/op-dl-server/define"
	"log"
	"strconv"
	"time"
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

// TimeFormat 字符串转时间
func TimeFormat(str string) time.Time {
	t, err := time.Parse(time.DateTime, str)
	if err != nil {
		panic(err)
	}
	return t
}
