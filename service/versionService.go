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

// UpdateVersion
// @Summary 更新版本信息
// @Description 更新版本信息
// @Tags 版本
// @param id formData int false "id"
// @param num formData int false "num"
// @param startTime formData string false "startTime"
// @param endTime formData string false "endTime"
// @param title formData string false "title"
// @param img formData string false "Img"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /version [put]
func UpdateVersion(c *gin.Context) {

	var version model.Version
	err := c.ShouldBind(&version)

	if err != nil {
		result.FailIllegalParameter(c)
		fmt.Println(err.Error())
		return
	}
	fmt.Println(version)
	err = util.DB.Model(&version).Omit("created_at").Save(&version).Error
	if err != nil {
		result.FailNormalError(c, "update error")
		return
	}
	result.Ok(c)
}

// CreateVersion
// @Summary 新增版本信息
// @Description 新增版本信息
// @Tags 版本
// @param num formData int false "num"
// @param startTime formData string false "startTime"
// @param endTime formData string false "endTime"
// @param title formData string false "title"
// @param img formData string false "Img"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /version [post]
func CreateVersion(c *gin.Context) {

	var version model.Version
	err := c.ShouldBind(&version)

	if err != nil {
		result.FailIllegalParameter(c)
		fmt.Println(err.Error())
		return
	}
	err = util.DB.Model(&version).Create(&version).Error
	if err != nil {
		result.FailNormalError(c, "create error")
		return
	}
	result.Ok(c)
}

// DeleteVersion
// @Summary 删除版本
// @Description 删除版本
// @Tags 版本
// @param id formData int false "id"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /version-delete [post]
func DeleteVersion(c *gin.Context) {

	var version model.Version
	err := c.ShouldBind(&version)

	if err != nil {
		result.FailIllegalParameter(c)
		fmt.Println(err.Error())
		return
	}
	fmt.Println(version)
	err = util.DB.Model(&version).Delete(&version).Error
	if err != nil {
		result.FailNormalError(c, "delete error")
		return
	}
	result.Ok(c)
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
