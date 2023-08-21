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
	result.OkWithData(c, define.DataWithTotal{Data: data, Total: count})
}

// GetPoolWithTag
// @Summary 获取卡池信息和标签
// @Description 获取卡池信息和标签
// @Tags 卡池
// @param page query int false "page"
// @param size query int false "size"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /pool-with-tag [get]
func GetPoolWithTag(c *gin.Context) {
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
	err := tx.
		Preload("PoolTag").
		Preload("PoolTag.Tag").
		Count(&count).
		Offset(page).
		Limit(size).
		Find(&data).Error
	if err != nil {
		result.SqlQueryError(c)
		return
	}
	result.OkWithData(c, define.DataWithTotal{Data: data, Total: count})
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

	err = util.DB.Model(&data).Omit("created_at").Save(&data).Error
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

// GetPoolTag
// @Summary 活动Tag
// @Description 获取Tag
// @Tags 卡池
// @param id query int false "id"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /pool-tag [get]
func GetPoolTag(c *gin.Context) {

	q := c.Query("id")

	if q == "" {
		result.FailIllegalParameter(c)
		return
	}

	var data []model.PoolTag
	var count int64
	err := util.DB.Model(&model.PoolTag{}).
		Where("pool_id=?", q).
		Count(&count).Find(&data).Error

	if err != nil {
		log.Println("database query error", err.Error())
	}
	result.OkWithData(c, define.DataWithTotal{Data: data, Total: count})
}

// CreatePoolTag
// @Summary 创建Tag
// @Description 创建Tag
// @Tags 卡池
// @param tagId query int false "tagId"
// @param eventId query int false "EventId"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /pool-with-tag [post]
func CreatePoolTag(c *gin.Context) {
	var data model.PoolTag
	err := c.ShouldBind(&data)

	if err != nil {
		result.FailIllegalParameter(c)
		fmt.Println(err.Error())
		return
	}

	err = util.DB.Model(&model.PoolTag{}).Create(&data).Error

	if err != nil {
		log.Println("database query error", err.Error())
	}
	result.Ok(c)

}

// DeletePoolTag
// @Summary 删除活Tag
// @Description 删除Tag
// @Tags 卡池
// @param tagId query int false "tagId"
// @param eventId query int false "eventId"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /pool-with-tag-delete [post]
func DeletePoolTag(c *gin.Context) {
	var data model.PoolTag
	err := c.ShouldBind(&data)

	fmt.Println(&data)
	if err != nil {
		result.FailIllegalParameter(c)
		fmt.Println(err.Error())
		return
	}
	util.DB.Model(&model.PoolTag{}).Where("tag_id=? AND pool_id=?", data.TagID, data.PoolId).Delete(&data)

	if err != nil {
		log.Println("database query error", err.Error())
	}
	result.Ok(c)
}
