package service

import (
	"github.com/gin-gonic/gin"
	"hasdsd.cn/op-dl-server/define"
	"hasdsd.cn/op-dl-server/model"
	"hasdsd.cn/op-dl-server/result"
	"hasdsd.cn/op-dl-server/util"
	"log"
)

// GetVersions
// @Summary 版本
// @Description 获取版本列表
// @Tags 版本
// @param page query int false "请输入当前页，默认第一页"
// @param size query int false "页大小"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /version [get]
func GetVersions(c *gin.Context) {

	page, size := util.GetPageInfo(c)

	var count int64
	data := make([]*model.Version, 0)
	err := util.DB.Model(&model.Version{}).Count(&count).Offset(page).Limit(size).Find(&data).Error

	if err != nil {
		log.Println("database query error")
	}
	result.OkWithData(c, define.DataWithTotal{Data: data, Total: count})
}

// GetVersionWithTag
// @Summary 版本和版本与Tag
// @Description 获取版本与Tag
// @Tags 版本
// @param page query int false "请输入当前页，默认第一页"
// @param size query int false "页大小"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /version-with-tag [get]
func GetVersionWithTag(c *gin.Context) {
	page, size := util.GetPageInfo(c)

	var data []model.Version
	var count int64
	err := util.DB.Model(&model.Version{}).
		Preload("VersionTag").
		Preload("VersionTag.Tag").
		Count(&count).Find(&data).
		Limit(page).Offset(size).Error

	if err != nil {
		log.Println("database query error", err.Error())
	}

	result.OkWithData(c, define.DataWithTotal{Data: data, Total: count})
}

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
