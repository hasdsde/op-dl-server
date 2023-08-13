package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hasdsd.cn/op-dl-server/model"
	"hasdsd.cn/op-dl-server/result"
	"hasdsd.cn/op-dl-server/util"
)

// GetUserVersionEvent
// @Summary 查询用户版本活动信息
// @Description 查询用户版本活动信息
// @Tags 用户版本活动
// @param userId query int false "userId"
// @param versionEventId query int false "versionEventId"
// @param page query int false "page"
// @param size query int false "size"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /user-version-event [get]
func GetUserVersionEvent(c *gin.Context) {

	userId := c.Query("userId")
	versionEventId := c.Query("versionEventId")

	page, size := util.GetPageInfo(c)
	data := make([]*model.UserVersionEvent, 0)

	var count int64
	tx := util.DB.Model(&model.UserVersionEvent{})

	if userId != "" {
		tx.Where("user_id=?", userId)
	}
	if versionEventId != "" {
		tx.Where("version_event_id=?", versionEventId)
	}
	err := tx.Count(&count).Offset(page).Limit(size).Find(&data).Error

	if err != nil {
		result.SqlQueryError(c)
		return
	}

	result.OkWithData(c, data)
}

// UpdateUserVersionEvent
// @Summary 更新用户版本活动信息
// @Description 更新用户版本活动信息
// @Tags 用户版本活动
// @param id formData int false "id"
// @param userId formData int false "userId"
// @param todoNum formData int false "todoNum"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /user-version-event [put]
func UpdateUserVersionEvent(c *gin.Context) {
	id := c.PostForm("id")
	userId := c.PostForm("userId")

	todo := c.PostForm("todoNum")
	if id == "" || todo == "" {
		result.FailIllegalParameter(c)
		return
	}
	err := util.DB.Model(&model.UserVersionEvent{}).Where("id=? and user_id=?", id, userId).Update("todo", todo).Error
	if err != nil {
		result.SqlQueryError(c)
		return
	}
	result.Ok(c)
}

// CreateUserVersionEvent
// @Summary 创建用户版本活动信息
// @Description 创建用户版本活动信息
// @Tags 用户版本活动
// @param userId formData int false "userId"
// @param versionEventId formData int false "versionEventId"
// @param todoNum formData int false "todoNum"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /user-version-event [post]
func CreateUserVersionEvent(c *gin.Context) {
	var data model.UserVersionEvent
	err := c.ShouldBind(&data)

	if err != nil {
		result.FailIllegalParameter(c)
		fmt.Println(err.Error())
		return
	}

	var count int64
	util.DB.Model(&data).Where("user_id = ? and version_event_id = ?", data.UserID, data.VersionEventID).Count(&count)
	if count > 0 {
		result.FailNormalError(c, "data already exist")
		return
	}

	err = util.DB.Model(&data).Create(&data).Error
	if err != nil {
		result.SqlQueryError(c)
		return
	}
	result.Ok(c)
}

// DeleteUserVersionEvent
// @Summary 删除用户版本活动信息
// @Description 删除用户版本活动信息
// @Tags 用户版本活动
// @param userId formData int false "userId"
// @param id formData int false "id"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /user-version-event-delete [post]
func DeleteUserVersionEvent(c *gin.Context) {

	var data model.UserVersionEvent
	err := c.ShouldBind(&data)

	if err != nil {
		result.FailIllegalParameter(c)
		fmt.Println(err.Error())
		return
	}

	err = util.DB.Model(&data).
		Where("user_id = ? and id = ?", data.UserID, data.ID).
		Delete(&data).
		Error

	if err != nil {
		result.SqlQueryError(c)
		return
	}
	result.Ok(c)
}
