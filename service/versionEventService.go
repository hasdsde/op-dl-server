package service

import (
	"github.com/gin-gonic/gin"
	"hasdsd.cn/op-dl-server/model"
	"hasdsd.cn/op-dl-server/result"
	"hasdsd.cn/op-dl-server/util"
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
