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
	err := tx.Count(&count).Order("id desc").Offset(page).Limit(size).Find(&data).Error
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
	userId := c.Query("userId")
	var user model.User
	//获取用户token
	err2 := util.DB.Model(&model.User{}).Where("user_id = ?", userId).First(&user).Error
	if err2 != nil {
		result.FailNormalError(c, "get user token error")
		return
	}
	//根据用户token获取每日详情
	noteInfo := util.GetDailyStatus(user.Token)
	var data model.Daily

	if noteInfo != nil {
		data.UserID = user.ID
		data.UID = noteInfo.GameInfo.GameUID
		data.HomeCoin = int64(noteInfo.HomeCoin.CurrentHomeCoin)
		data.TaskNum = int64(noteInfo.Task.FinishedTaskNum)
		if noteInfo.Task.IsExtraTaskRewardReceived {
			data.TaskFinished = 1
		} else {
			data.TaskFinished = 0
		}
		data.Resin = int64(noteInfo.Resin.CurrentResin)
		data.Expeditions = int64(noteInfo.Expeditions.CurrentExpeditionNum)
		if noteInfo.Transformer.RecoveryTime.Reached {
			data.Transformer = 1
		} else {
			data.Transformer = 0
		}
		data.ResinDiscount = int64(noteInfo.ResinDiscount.ResinUnusedDiscountNum)
	} else {
		result.FailNormalError(c, "get user daily info error")
		return
	}

	err := util.DB.Model(&model.Daily{}).Create(data).Error
	if err != nil {
		result.SqlQueryError(c)
		return
	}
}
