package service

import (
	"github.com/gin-gonic/gin"
	"hasdsd.cn/op-dl-server/define"
	"hasdsd.cn/op-dl-server/result"
	"hasdsd.cn/op-dl-server/util"
	"log"
)

var v = util.DB.Version
var t = util.DB.Tag
var vt = util.DB.VersionTag

// GetVersions
// @Summary 版本
// @Description 获取版本列表
// @Tags 版本
// @param page query int false "请输入当前页，默认第一页"
// @param size query int false "页大小"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /version [get]
func GetVersions(ctx *gin.Context) {

	page, size := util.GetPageInfo(ctx)

	data, count, err := v.WithContext(ctx).Order(v.ID.Desc()).FindByPage(page, size)
	if err != nil {
		return
	}
	if err != nil {
		log.Println("database query error")
	}
	result.OkWithData(ctx, define.DataWithTotal{Data: data, Total: count})
}

// GetVersionWithTag
// @Summary 版本和版本详情
// @Description 获取版本与版本详情
// @Tags 版本
// @param page query int false "请输入当前页，默认第一页"
// @param size query int false "页大小"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /version-with-tag [get]
func GetVersionWithTag(ctx *gin.Context) {
	page, size := util.GetPageInfo(ctx)
	//data := make([]*define.VersionDetails, 0)

	data, count, err := v.WithContext(ctx).Select(v.ALL).FindByPage(page, size)

	if err != nil {
		log.Println("database query error")
	}

	result.OkWithData(ctx, define.DataWithTotal{Data: data, Total: count})
}
