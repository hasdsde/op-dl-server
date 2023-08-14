package service

import (
	"github.com/gin-gonic/gin"
	"hasdsd.cn/op-dl-server/define"
	"hasdsd.cn/op-dl-server/model"
	"hasdsd.cn/op-dl-server/result"
	"hasdsd.cn/op-dl-server/util"
	"log"
)

// GetUserDaily
// @Summary 每日任务
// @Description 获取用户每日任务
// @Tags 每日任务
// @param page query int false "请输入当前页，默认第一页"
// @param size query int false "页大小"
// @param userId query int false "userId"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /daily [get]
func GetUserDaily(c *gin.Context) {
	page, size := util.GetPageInfo(c)
	userId := c.Query("userId")

	var data []*model.Daily
	var count int64
	tx := util.DB.Model(&model.Daily{})
	if userId != "" {
		tx.Where("user_id = ?", userId)
	}
	err := tx.Count(&count).Find(&data).Offset(page).Limit(size).Error
	if err != nil {
		log.Println("database query error")
	}
	result.OkWithData(c, define.DataWithTotal{Data: data, Total: count})
}

// FreshUserDaily
// @Summary 刷新每日任务
// @Description 刷新用户每日任务
// @Tags 每日任务
// @param userId query int false "userId"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /daily [post]
func FreshUserDaily(c *gin.Context) {

}
