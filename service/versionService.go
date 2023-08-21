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
		fmt.Println("Error: " + err.Error())
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

// GetVersionTag
// @Summary 版本Tag
// @Description 获取Tag
// @Tags 版本
// @param id query int false "id"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /version-tag [get]
func GetVersionTag(c *gin.Context) {

	versionId := c.Query("id")

	if versionId == "" {
		result.FailIllegalParameter(c)
		return
	}

	var data []model.VersionTag
	var count int64
	err := util.DB.Model(&model.VersionTag{}).
		Where("version_id=?", versionId).
		Count(&count).Find(&data).Error

	if err != nil {
		log.Println("database query error", err.Error())
	}
	result.OkWithData(c, define.DataWithTotal{Data: data, Total: count})
}

// DeleteVersionTag
// @Summary 删除版本和版本与Tag
// @Description 删除版本与Tag
// @Tags 版本
// @param tagId query int false "tagId"
// @param versionId query int false "versionId"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /version-tag-delete [post]
func DeleteVersionTag(c *gin.Context) {
	var versionTag model.VersionTag
	err := c.ShouldBind(&versionTag)

	fmt.Println(&versionTag)
	if err != nil {
		result.FailIllegalParameter(c)
		fmt.Println(err.Error())
		return
	}
	util.DB.Model(&model.VersionTag{}).Where("tag_id=? AND version_id=?", versionTag.TagID, versionTag.VersionID).Delete(&versionTag)

	if err != nil {
		log.Println("database query error", err.Error())
	}
	result.Ok(c)
}

// CreateVersionTag
// @Summary 创建版本与Tag
// @Description 创建版本与Tag
// @Tags 版本
// @param tagId query int false "tagId"
// @param versionId query int false "versionId"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /version-tag [post]
func CreateVersionTag(c *gin.Context) {
	var data model.VersionTag
	err := c.ShouldBind(&data)

	if err != nil {
		result.FailIllegalParameter(c)
		fmt.Println(err.Error())
		return
	}

	err = util.DB.Model(&model.VersionTag{}).Create(&data).Error

	if err != nil {
		log.Println("database query error", err.Error())
	}
	result.Ok(c)
}
