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

// GetTag
// @Summary 标签
// @Description 获取标签列表
// @Tags 标签
// @param page query int false "请输入当前页，默认第一页"
// @param sort query string false "分类"
// @param size query int false "页大小"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /tag [get]
func GetTag(c *gin.Context) {
	page, size := util.GetPageInfo(c)
	var count int64
	data := make([]*model.Tag, 0)
	sort := c.Query("sort")

	tx := util.DB.Model(&model.Tag{})
	if sort == "全部" {
		sort = ""
	}
	if sort != "" {
		tx.Where("sort = ?", sort)
	}
	err := tx.
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

// CreateTag
// @Summary 新增标签信息
// @Description 新增标签信息
// @Tags 标签
// @param title formData string false "title"
// @param primogems formData int false "primogems"
// @param award formData string false "award"
// @param preTask formData int false "preTask"
// @param coPlay formData int false "coPlay"
// @param todoNum formData int false "todoNum"
// @param startTime formData string false "startTime"
// @param endTime formData string false "endTime"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /tag [post]
func CreateTag(c *gin.Context) {
	var data model.Tag
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

// UpdateTag
// @Summary 更新标签信息
// @Description 更新标签信息
// @Tags 标签
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
// @Router /tag [put]
func UpdateTag(c *gin.Context) {
	var data model.Tag
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

// DeleteTag
// @Summary 删除标签
// @Description 删除标签
// @Tags 标签
// @param id formData int false "id"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /tag-delete [post]
func DeleteTag(c *gin.Context) {
	var data model.Tag
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
