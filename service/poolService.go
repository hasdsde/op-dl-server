package service

import (
	"github.com/gin-gonic/gin"
	"hasdsd.cn/op-dl-server/model"
	"hasdsd.cn/op-dl-server/result"
	"hasdsd.cn/op-dl-server/util"
)

// GetPool
// @Summary 获取卡池信息
// @Description 获取卡池信息
// @Tags 卡池
// @param page query int false "page"
// @param size query int false "size"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /pool [get]
func GetPool(c *gin.Context) {
	page, size := util.GetPageInfo(c)
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	var data []*model.Pool
	var count int64
	tx := util.DB.Model(&model.Pool{})
	if startTime != "" {
		tx.Where("start_time < ?", startTime)
	}
	if endTime != "" {
		tx.Where("end_time > ?", endTime)
	}
	err := tx.Count(&count).Offset(page).Limit(size).Find(&data).Error
	if err != nil {
		result.SqlQueryError(c)
		return
	}
	result.OkWithData(c, data)
}

// UpdatePool
// @Summary 更新卡池
// @Description 获取卡池信息
// @Tags 卡池
// @param page query int false "page"
// @param size query int false "size"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /pool [put]
func UpdatePool(c *gin.Context) {
	var data model.Pool
	err := c.ShouldBind(&data)
	if err != nil {
		result.FailIllegalParameter(c)
		return
	}

	err = util.DB.Model(&data).Save(&data).Error
	if err != nil {
		result.SqlQueryError(c)
		return
	}
	result.Ok(c)
}

// CreatePool
// @Summary 新增卡池
// @Description 新增卡池信息
// @Tags 卡池
// @param versionNum query int false "versionNum"
// @param name query string false "name"
// @param Img query string false "Img"
// @param Star query int false "Star"
// @param Type query string false "Type"
// @param StartTime query int false "StartTime"
// @param EndTime query int false "EndTime"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /pool [post]
func CreatePool(c *gin.Context) {
	var data model.Pool
	err := c.ShouldBind(&data)
	if err != nil {
		result.FailIllegalParameter(c)
		return
	}

	err = util.DB.Model(&data).Create(&data).Error
	if err != nil {
		result.SqlQueryError(c)
		return
	}
	result.Ok(c)
}

// DelPool
// @Summary 删除卡池
// @Description 删除卡池
// @Tags 卡池
// @param userId formData int false "userId"
// @param id formData int false "id"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /pool-delete [post]
func DelPool(c *gin.Context) {
	var data model.Pool
	err := c.ShouldBind(&data)
	if err != nil {
		result.FailIllegalParameter(c)
		return
	}
	err = util.DB.Delete(&data).Error
	if err != nil {
		result.SqlQueryError(c)
		return
	}
	result.Ok(c)
}
