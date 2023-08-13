package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hasdsd.cn/op-dl-server/model"
	"hasdsd.cn/op-dl-server/result"
	"hasdsd.cn/op-dl-server/util"
)

// GetUserEvent
// @Summary 查询用户活动信息
// @Description 查询用户活动信息
// @Tags 用户活动
// @param userId query int false "userId"
// @param eventId query int false "eventId"
// @param page query int false "page"
// @param size query int false "size"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /user-event [get]
func GetUserEvent(c *gin.Context) {

	userId := c.Query("userId")
	eventId := c.Query("eventId")

	page, size := util.GetPageInfo(c)
	data := make([]*model.UserEvent, 0)

	var count int64
	tx := util.DB.Model(&model.UserEvent{})

	if userId != "" {
		tx.Where("user_id=?", userId)
	}
	if eventId != "" {
		tx.Where("event_id=?", eventId)
	}
	err := tx.Count(&count).Offset(page).Limit(size).Find(&data).Error

	if err != nil {
		result.SqlQueryError(c)
		return
	}

	result.OkWithData(c, data)
}

// UpdateUserEvent
// @Summary 更新用户活动信息
// @Description 更新用户活动信息
// @Tags 用户活动
// @param id formData int false "id"
// @param userId formData int false "userId"
// @param todoNum formData int false "todoNum"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /user-event [put]
func UpdateUserEvent(c *gin.Context) {
	id := c.PostForm("id")
	userId := c.PostForm("userId")

	todo := c.PostForm("todoNum")
	if id == "" || todo == "" {
		result.FailIllegalParameter(c)
		return
	}
	err := util.DB.Model(&model.UserEvent{}).Where("id=? and user_id=?", id, userId).Update("todo", todo).Error
	if err != nil {
		result.SqlQueryError(c)
		return
	}
	result.Ok(c)
}

// CreateUserEvent
// @Summary 创建用户活动信息
// @Description 创建用户活动信息
// @Tags 用户活动
// @param userId formData int false "userId"
// @param eventId formData int false "eventId"
// @param todoNum formData int false "todoNum"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /user-event [post]
func CreateUserEvent(c *gin.Context) {
	var data model.UserEvent
	err := c.ShouldBind(&data)

	if err != nil {
		result.FailIllegalParameter(c)
		fmt.Println(err.Error())
		return
	}

	var count int64
	util.DB.Model(&data).Where("user_id = ? and event_id = ?", data.UserID, data.EventID).Count(&count)
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

// DeleteUserEvent
// @Summary 删除用户活动信息
// @Description 删除用户活动信息
// @Tags 用户活动
// @param userId formData int false "userId"
// @param id formData int false "id"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /user-event-delete [post]
func DeleteUserEvent(c *gin.Context) {

	var data model.UserEvent
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
