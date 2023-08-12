package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hasdsd.cn/op-dl-server/model"
	"hasdsd.cn/op-dl-server/result"
	"hasdsd.cn/op-dl-server/util"
	"log"
)

// GetVersionEvent
// @Summary 根据版本获取版本活动
// @Description 获取版本活动
// @Tags 版本活动
// @param version_num query int false "version_num"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /version-event [get]
func GetVersionEvent(c *gin.Context) {
	versionNum := c.Query("version_num")
	if versionNum == "" {
		result.FailIllegalParameter(c)
		return
	}
	var data []*model.VersionEvent
	err := util.DB.
		Model(&model.VersionEvent{}).
		Where("version_num=?", versionNum).
		Find(&data).Error

	if err != nil {
		result.FailNormalError(c, "sql error"+err.Error())
		return
	}
	result.OkWithData(c, data)
}

// CreateVersionEvent
// @Summary 新增版本活动
// @Description 新增版本活动
// @Tags 版本活动
// @param versionNum formData int false "versionNum"
// @param title formData string false "title"
// @param content formData string false "content"
// @param level formData int false "level"
// @param startTime formData string false "startTime"
// @param endTime formData string false "endTime"
// @param todoNum formData string false "todoNum"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /version-event [post]
func CreateVersionEvent(c *gin.Context) {
	var versionEvent model.VersionEvent
	err := c.ShouldBind(&versionEvent)
	if err != nil {
		result.FailIllegalParameter(c)
		return
	}
	err = util.DB.Model(&model.VersionEvent{}).Create(&versionEvent).Error
	if err != nil {
		result.FailNormalError(c, "sql error"+err.Error())
		return
	}
	result.Ok(c)
}

// UpdateVersionEvent
// @Summary 修改版本活动
// @Description 修改版本活动
// @Tags 版本活动
// @param id formData int false "id"
// @param versionNum formData int false "versionNum"
// @param title formData string false "title"
// @param content formData string false "content"
// @param level formData int false "level"
// @param startTime formData string false "startTime"
// @param endTime formData string false "endTime"
// @param todoNum formData string false "todoNum"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /version-event [put]
func UpdateVersionEvent(c *gin.Context) {
	var versionEvent model.VersionEvent
	err := c.ShouldBind(&versionEvent)
	if err != nil {
		result.FailIllegalParameter(c)
		fmt.Println(err.Error())
		return
	}
	fmt.Println(versionEvent)
	err = util.DB.Model(&versionEvent).Omit("created_at").Save(&versionEvent).Error
	if err != nil {
		result.FailNormalError(c, "update error")
		return
	}
	result.Ok(c)
}

// DeleteVersionEvent
// @Summary 删除版本活动
// @Description 删除版本活动
// @Tags 版本活动
// @param id formData int false "id"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /version-event-delete [post]
func DeleteVersionEvent(c *gin.Context) {
	var data model.VersionEvent
	err := c.ShouldBind(&data)
	if err != nil {
		result.FailIllegalParameter(c)
		fmt.Println(err.Error())
		return
	}
	fmt.Println(data)
	err = util.DB.Model(&data).Delete(&data).Error
	if err != nil {
		result.FailNormalError(c, "delete error")
		return
	}
	result.Ok(c)
}

// GetVersionEventWithTag
// @Summary 根据版本获取版本活动和tag
// @Description 获取版本活动和tag
// @Tags 版本活动
// @param version_num query int false "version_num"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /version-event-with-tag [get]
func GetVersionEventWithTag(c *gin.Context) {
	versionNum := c.Query("version_num")
	if versionNum == "" {
		result.FailIllegalParameter(c)
		return
	}
	var data []*model.VersionEvent
	err := util.DB.
		Model(&model.VersionEvent{}).
		Where("version_num=?", versionNum).
		Preload("VersionEventTag").
		Preload("VersionEventTag.Tag").
		Find(&data).Error

	if err != nil {
		result.FailNormalError(c, "sql error"+err.Error())
		return
	}
	result.OkWithData(c, data)
}

// DeleteVersionEventTag
// @Summary 删除版本活动与Tag
// @Description 删除版本活动与Tag
// @Tags 版本活动
// @param tagId query int false "tagId"
// @param versionEventId query int false "versionEventId"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /version-event-with-tag-delete [post]
func DeleteVersionEventTag(c *gin.Context) {
	var data model.VersionEventTag
	err := c.ShouldBind(&data)

	if err != nil {
		result.FailIllegalParameter(c)
		fmt.Println(err.Error())
		return
	}
	util.DB.Model(&model.VersionEventTag{}).Where("tag_id=? AND version_event_id=?", data.TagID, data.VersionEventID).Delete(&data)

	if err != nil {
		log.Println("database query error", err.Error())
	}
	result.Ok(c)
}
