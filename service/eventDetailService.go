package service

import (
	"github.com/gin-gonic/gin"
	"hasdsd.cn/op-dl-server/model"
	"hasdsd.cn/op-dl-server/result"
	"hasdsd.cn/op-dl-server/util"
)

// GetEventDetail
// @Summary 获取活动细节信息
// @Description 获取活动细节信息
// @Tags 活动细节
// @param page query int false "page"
// @param size query int false "size"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /eventDetail [get]
func GetEventDetail(c *gin.Context) {
	page, size := util.GetPageInfo(c)
	eventId := c.Query("eventId")
	var data []*model.EventDetail
	var count int64
	tx := util.DB.Model(&model.EventDetail{})
	if eventId != "" {
		tx.Where("event_id = ?", eventId)
	}
	err := tx.Count(&count).Offset(page).Limit(size).Find(&data).Error
	if err != nil {
		result.SqlQueryError(c)
		return
	}
	result.OkWithData(c, data)
}

// UpdateEventDetail
// @Summary 更新活动细节
// @Description 获取活动细节信息
// @Tags 活动细节
// @param page query int false "page"
// @param size query int false "size"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /eventDetail [put]
func UpdateEventDetail(c *gin.Context) {
	var data model.EventDetail
	err := c.ShouldBind(&data)
	if err != nil {
		result.FailIllegalParameter(c)
		return
	}

	err = util.DB.Model(&data).Save(&data).Error
	if err != nil {
		result.SqlQueryError(c)
		return
	}
	result.Ok(c)
}

// CreateEventDetail
// @Summary 新增活动细节
// @Description 新增活动细节信息
// @Tags 活动细节
// @param versionNum query int false "versionNum"
// @param name query string false "name"
// @param Img query string false "Img"
// @param Star query int false "Star"
// @param Type query string false "Type"
// @param StartTime query int false "StartTime"
// @param EndTime query int false "EndTime"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /eventDetail [post]
func CreateEventDetail(c *gin.Context) {
	var data model.EventDetail
	err := c.ShouldBind(&data)
	if err != nil {
		result.FailIllegalParameter(c)
		return
	}

	err = util.DB.Model(&data).Create(&data).Error
	if err != nil {
		result.SqlQueryError(c)
		return
	}
	result.Ok(c)
}

// DelEventDetail
// @Summary 删除活动细节
// @Description 删除活动细节
// @Tags 活动细节
// @param userId formData int false "userId"
// @param id formData int false "id"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /eventDetail-delete [post]
func DelEventDetail(c *gin.Context) {
	var data model.EventDetail
	err := c.ShouldBind(&data)
	if err != nil {
		result.FailIllegalParameter(c)
		return
	}
	err = util.DB.Delete(&data).Error
	if err != nil {
		result.SqlQueryError(c)
		return
	}
	result.Ok(c)
}
