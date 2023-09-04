package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hasdsd.cn/op-dl-server/define"
	"hasdsd.cn/op-dl-server/model"
	"hasdsd.cn/op-dl-server/result"
	"hasdsd.cn/op-dl-server/util"
	"log"
)

// GetEvent
// @Summary 活动
// @Description 获取活动列表
// @Tags 活动
// @param page query int false "请输入当前页，默认第一页"
// @param size query int false "页大小"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /event [get]
func GetEvent(c *gin.Context) {
	page, size := util.GetPageInfo(c)
	var count int64
	data := make([]*model.Event, 0)

	err := util.DB.Model(&model.Event{}).
		Count(&count).
		Offset(page).
		Limit(size).
		Find(&data).
		Error

	if err != nil {
		log.Println("database query error")
	}
	result.OkWithData(c, define.DataWithTotal{Data: data, Total: count})
}

// CreateEvent
// @Summary 新增活动信息
// @Description 新增活动信息
// @Tags 活动
// @param title formData string false "title"
// @param primogems formData int false "primogems"
// @param award formData string false "award"
// @param preTask formData int false "preTask"
// @param coPlay formData int false "coPlay"
// @param todoNum formData int false "todoNum"
// @param startTime formData string false "startTime"
// @param endTime formData string false "endTime"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /event [post]
func CreateEvent(c *gin.Context) {
	var data model.Event
	err := c.ShouldBind(&data)
	if err != nil {
		result.FailIllegalParameter(c)
		fmt.Println(err.Error())
		return
	}
	err = util.DB.Model(&data).Create(&data).Error
	if err != nil {
		result.FailNormalError(c, "create error")
		return
	}
	result.Ok(c)
}

// UpdateEvent
// @Summary 更新活动信息
// @Description 更新活动信息
// @Tags 活动
// @param id formData int false "id"
// @param title formData string false "title"
// @param primogems formData int false "primogems"
// @param award formData string false "award"
// @param preTask formData int false "preTask"
// @param coPlay formData int false "coPlay"
// @param todoNum formData int false "todoNum"
// @param startTime formData string false "startTime"
// @param endTime formData string false "endTime"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /event [put]
func UpdateEvent(c *gin.Context) {
	var data model.Event
	err := c.ShouldBind(&data)
	if err != nil {
		result.FailIllegalParameter(c)
		fmt.Println(err.Error())
		return
	}
	fmt.Println(data)
	err = util.DB.Model(&data).Omit("created_at").Save(&data).Error
	if err != nil {
		result.FailNormalError(c, "update error")
		return
	}
	result.Ok(c)
}

// DeleteEvent
// @Summary 删除活动
// @Description 删除活动
// @Tags 活动
// @param id formData int false "id"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /event-delete [post]
func DeleteEvent(c *gin.Context) {
	var data model.Event
	err := c.ShouldBind(&data)
	if err != nil {
		result.FailIllegalParameter(c)
		fmt.Println(err.Error())
		return
	}
	err = util.DB.Model(&data).Delete(&data).Error
	if err != nil {
		result.FailNormalError(c, "delete error")
		return
	}
	result.Ok(c)
}

// GetEventWithTag
// @Summary 活动和活动与Tag
// @Description 获取活动与Tag
// @Tags 活动
// @param page query int false "请输入当前页，默认第一页"
// @param size query int false "页大小"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /event-with-tag [get]
func GetEventWithTag(c *gin.Context) {
	page, size := util.GetPageInfo(c)

	var data []model.Event
	var count int64
	err := util.DB.Model(&model.Event{}).
		Preload("EventTag").
		Preload("EventTag.Tag").
		Count(&count).Find(&data).
		Limit(page).Offset(size).Error

	if err != nil {
		log.Println("database query error", err.Error())
	}
	result.OkWithData(c, define.DataWithTotal{Data: data, Total: count})
}

// GetEventWithTagWithEventDetail
// @Summary 活动和活动与Tag和详情
// @Description 获取活动与Tag和详情
// @Tags 活动
// @param page query int false "请输入当前页，默认第一页"
// @param size query int false "页大小"
// @param time query int false "time"
// @param versionId query int false "versionId"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /event-with-tag-detail [get]
func GetEventWithTagWithEventDetail(c *gin.Context) {
	page, size := util.GetPageInfo(c)

	var data []model.Event
	var count int64
	tx := util.DB.Model(&model.Event{})

	versionId := c.Query("versionId")
	//时间筛选
	tx = util.QueryWithTime(tx, c)

	//外键和分页
	err := tx.Where("version_num = ?", versionId).
		Preload("EventTag").
		Preload("EventTag.Tag").
		Preload("EventDetail").
		Count(&count).Find(&data).
		Limit(page).Offset(size).Error

	if err != nil {
		log.Println("database query error", err.Error())
	}
	result.OkWithData(c, define.DataWithTotal{Data: data, Total: count})
}

// GetCurrentEventWithTagWithEventDetail
// @Summary 当前活动和活动与Tag和详情
// @Description 当前获取活动与Tag和详情
// @Tags 活动
// @param page query int false "请输入当前页，默认第一页"
// @param size query int false "页大小"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /current-event-with-tag-detail [get]
func GetCurrentEventWithTagWithEventDetail(c *gin.Context) {
	page, size := util.GetPageInfo(c)

	var data []model.Event
	var count int64
	err := util.DB.Model(&model.Event{}).
		Where("end_time > ? and start_time < ?", util.GetLocalTime(), util.GetLocalTime()).
		Preload("EventTag").
		Preload("EventTag.Tag").
		Preload("EventDetail").
		Count(&count).Find(&data).
		Limit(page).Offset(size).Error

	if err != nil {
		log.Println("database query error", err.Error())
	}
	result.OkWithData(c, define.DataWithTotal{Data: data, Total: count})
}

// GetEventTag
// @Summary 活动Tag
// @Description 获取Tag
// @Tags 活动
// @param id query int false "id"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /event-tag [get]
func GetEventTag(c *gin.Context) {

	q := c.Query("id")

	if q == "" {
		result.FailIllegalParameter(c)
		return
	}

	var data []model.EventTag
	var count int64
	err := util.DB.Model(&model.EventTag{}).
		Where("event_id=?", q).
		Count(&count).Find(&data).Error

	if err != nil {
		log.Println("database query error", err.Error())
	}
	result.OkWithData(c, define.DataWithTotal{Data: data, Total: count})
}

// CreateEventTag
// @Summary 创建活动与Tag
// @Description 创建活动与Tag
// @Tags 活动
// @param tagId query int false "tagId"
// @param eventId query int false "EventId"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /event-with-tag [post]
func CreateEventTag(c *gin.Context) {
	var data model.EventTag
	err := c.ShouldBind(&data)

	if err != nil {
		result.FailIllegalParameter(c)
		fmt.Println(err.Error())
		return
	}

	err = util.DB.Model(&model.EventTag{}).Create(&data).Error

	if err != nil {
		log.Println("database query error", err.Error())
	}
	result.Ok(c)

}

// DeleteEventTag
// @Summary 删除活动和活动与Tag
// @Description 删除活动与Tag
// @Tags 活动
// @param tagId query int false "tagId"
// @param eventId query int false "eventId"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /event-with-tag-delete [post]
func DeleteEventTag(c *gin.Context) {
	var data model.EventTag
	err := c.ShouldBind(&data)

	if err != nil {
		result.FailIllegalParameter(c)
		fmt.Println(err.Error())
		return
	}
	util.DB.Model(&model.EventTag{}).Where("tag_id=? AND event_id=?", data.TagID, data.EventID).Delete(&data)

	if err != nil {
		log.Println("database query error", err.Error())
	}
	result.Ok(c)
}
