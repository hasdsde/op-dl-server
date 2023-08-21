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

// GetVersionEvent
// @Summary 根据版本获取版本活动
// @Description 获取版本活动
// @Tags 版本活动
// @param page query int false "请输入当前页，默认第一页"
// @param size query int false "页大小"
// @param version_num query int false "version_num"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /version-event [get]
func GetVersionEvent(c *gin.Context) {
	page, size := util.GetPageInfo(c)
	versionNum := c.Query("version_num")
	var data []*model.VersionEvent
	var count int64
	tx := util.DB.Model(&model.VersionEvent{})
	err := tx
	if versionNum != "" {
		tx.Where("version_num=?", versionNum)
	}
	err2 := tx.Count(&count).Offset(page).Limit(size).Find(&data).Error

	if err != nil {
		result.FailNormalError(c, "sql error"+err2.Error())
		return
	}
	result.OkWithData(c, define.DataWithTotal{Data: data, Total: count})
}

// CreateVersionEvent
// @Summary 新增活动
// @Description 新增活动
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
// @param page query int false "请输入当前页，默认第一页"
// @param size query int false "页大小"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /version-event-with-tag [get]
func GetVersionEventWithTag(c *gin.Context) {
	page, size := util.GetPageInfo(c)
	versionNum := c.Query("version_num")
	var data []*model.VersionEvent
	var count int64
	tx := util.DB.Model(&model.VersionEvent{})

	if versionNum != "" {
		tx.Where("version_num=?", versionNum)
	}

	err2 := tx.Preload("VersionEventTag").
		Preload("VersionEventTag.Tag").
		Count(&count).Offset(page).Limit(size).Find(&data).Error

	if err2 != nil {
		result.FailNormalError(c, "sql error"+err2.Error())
		return
	}
	result.OkWithData(c, define.DataWithTotal{Data: data, Total: count})

}

// CreateVersionEventTag
// @Summary 创建版本活动与Tag
// @Description 创建版本活动与Tag
// @Tags 版本活动
// @param tagId query int false "tagId"
// @param versionEventId query int false "versionEventId"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /version-event-with-tag [post]
func CreateVersionEventTag(c *gin.Context) {
	var data model.VersionEventTag
	err := c.ShouldBind(&data)

	if err != nil {
		result.FailIllegalParameter(c)
		fmt.Println(err.Error())
		return
	}

	err = util.DB.Model(&model.VersionEventTag{}).Create(&data).Error

	if err != nil {
		log.Println("database query error", err.Error())
	}
	result.Ok(c)
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
	err = util.DB.Model(&model.VersionEventTag{}).Where("tag_id=? AND version_event_id=?", data.TagID, data.VersionEventID).Delete(&data).Error

	if err != nil {
		log.Println("database query error", err.Error())
	}
	result.Ok(c)
}
