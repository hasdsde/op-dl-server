package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hasdsd.cn/op-dl-server/model"
	"hasdsd.cn/op-dl-server/result"
	"hasdsd.cn/op-dl-server/util"
	"time"
)

// GetUserPrivate
// @Summary 查询用户自定义活动信息
// @Description 查询用户自定义活动信息
// @Tags 用户自定义活动
// @param userId query int false "userId"
// @param page query int false "page"
// @param size query int false "size"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /user-private [get]
func GetUserPrivate(c *gin.Context) {

	userId := c.Query("userId")

	page, size := util.GetPageInfo(c)
	data := make([]*model.UserPrivate, 0)

	var count int64
	tx := util.DB.Model(&model.UserPrivate{})

	if userId != "" {
		tx.Where("user_id=?", userId)
	}

	err := tx.Count(&count).Offset(page).Limit(size).Find(&data).Error

	if err != nil {
		result.SqlQueryError(c)
		return
	}

	result.OkWithData(c, data)
}

// UpdateUserPrivate
// @Summary 更新用户自定义活动信息
// @Description 更新用户自定义活动信息
// @Tags 用户自定义活动
// @param id formData int false "id"
// @param userId formData int false "userId"
// @param todoNum formData int false "todoNum"
// @param todo formData int false "todo"
// @param startTime formData string false "startTime"
// @param endTime formData string false "endTime"
// @param autoUpdate formData string false "autoUpdate"
// @param updateHour formData string false "updateHour"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /user-private [put]
func UpdateUserPrivate(c *gin.Context) {
	var data model.UserPrivate
	err := c.ShouldBind(&data)
	if err != nil {
		result.FailIllegalParameter(c)
		return
	}
	if data.ID == 0 {
		result.FailIllegalParameter(c)
		return
	}
	tx := util.DB.Model(&model.UserPrivate{}).Where("id=? and user_id=?", data.ID, data.UserID)
	if data.Todo != 0 {
		tx.Update("todo", data.Todo)
	}
	if data.TodoNum != 0 {
		tx.Update("todo_num", data.TodoNum)
	}
	nilTime := time.Time{}
	if data.StartTime != nilTime {
		tx.Update("start_time", data.StartTime)
	}
	if data.EndTime != nilTime {
		tx.Update("end_time", data.EndTime)
	}
	if data.AutoUpdate != "" {
		tx.Update("auto_update", data.AutoUpdate)
	}
	if data.UpdateHour != 0 {
		tx.Update("update_hour", data.UpdateHour)
	}

	err = tx.Error
	if err != nil {
		result.SqlQueryError(c)
		return
	}
	result.Ok(c)
}

// CreateUserPrivate
// @Summary 创建用户自定义活动信息
// @Description 创建用户自定义活动信息
// @Tags 用户自定义活动
// @param userId formData int false "userId"
// @param todoNum formData int false "todoNum"
// @param todo formData int false "todo"
// @param startTime formData string false "startTime"
// @param endTime formData string false "endTime"
// @param autoUpdate formData string false "autoUpdate"
// @param updateHour formData string false "updateHour"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /user-private [post]
func CreateUserPrivate(c *gin.Context) {
	var data model.UserPrivate
	err := c.ShouldBind(&data)

	if err != nil {
		result.FailIllegalParameter(c)
		fmt.Println(err.Error())
		return
	}

	err = util.DB.Model(&data).Create(&data).Error
	if err != nil {
		result.SqlQueryError(c)
		return
	}
	result.Ok(c)
}

// DeleteUserPrivate
// @Summary 删除用户自定义活动信息
// @Description 删除用户自定义活动信息
// @Tags 用户自定义活动
// @param userId formData int false "userId"
// @param id formData int false "id"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /user-private-delete [post]
func DeleteUserPrivate(c *gin.Context) {

	var data model.UserPrivate
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
