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

// Login 登录
// @Summary 登录
// @Description 登录
// @Tags 用户
// @param name formData string false "name"
// @param password formData string false "name"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /login [post]
func Login(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")

	if name == "" || password == "" {
		result.FailIllegalParameter(c)
		return
	}
	var data model.User
	err := util.DB.Model(&model.User{}).
		Where("name=? and password =?", name, util.GenMd5(password)).
		First(&data).Error
	if err != nil {
		result.FailNormalError(c, "name or password error!")
		return
	}

	token, err := util.GenerateToken(data.ID, data.Name, data.Password)
	if err != nil {
		result.FailNormalError(c, "token generate error: "+err.Error())
		return
	}

	result.OkWithData(c, token)
}

// Register
// @Summary 注册
// @Description 注册
// @Tags 用户
// @param name formData string false "name"
// @param password formData string false "password"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /register [post]
func Register(c *gin.Context) {
	var data model.User
	err := c.ShouldBind(&data)

	if err != nil {
		result.FailIllegalParameter(c)
		fmt.Println(err.Error())
		return
	}

	data.Role = "user"
	data.Password = util.GenMd5(data.Password)
	err = util.DB.Model(&data).Create(&data).Error
	if err != nil {
		result.FailNormalError(c, "create error: name has been used !")
		return
	}
	result.Ok(c)
}

// GetUser
// @Summary 获取用户列表
// @Description 获取用户列表
// @Tags 用户
// @param page query int false "请输入当前页，默认第一页"
// @param size query int false "页大小"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /user [get]
func GetUser(c *gin.Context) {
	page, size := util.GetPageInfo(c)

	var data = make([]*model.User, 0)
	var count int64
	err := util.DB.Model(&model.User{}).
		Count(&count).
		Offset(page).
		Limit(size).
		Omit("password").
		Find(&data).
		Error
	if err != nil {
		log.Println("database query error")
	}
	result.OkWithData(c, define.DataWithTotal{Data: data, Total: count})

}

// CreateUser
// @Summary 新增用户信息
// @Description 新增用户信息
// @Tags 用户
// @param name formData string false "name"
// @param password formData string false "password"
// @param token formData string false "token"
// @param role formData string false "role"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /user [post]
func CreateUser(c *gin.Context) {
	var data model.User
	err := c.ShouldBind(&data)
	if err != nil {
		result.FailIllegalParameter(c)
		fmt.Println(err.Error())
		return
	}
	data.Password = util.GenMd5(data.Password)
	err = util.DB.Model(&data).Create(&data).Error
	if err != nil {
		result.FailNormalError(c, "create error")
		return
	}
	result.Ok(c)
}

// UpdateUser
// @Summary 修改用户信息
// @Description 修改用户信息
// @Tags 用户
// @param id formData int false "id"
// @param name formData string false "name"
// @param token formData string false "token"
// @param role formData string false "role"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /user [put]
func UpdateUser(c *gin.Context) {
	var data model.User
	err := c.ShouldBind(&data)
	if err != nil {
		result.FailIllegalParameter(c)
		fmt.Println(err.Error())
		return
	}
	fmt.Println(data)
	err = util.DB.Model(&data).Omit("created_at", "password").Save(&data).Error
	if err != nil {
		result.FailNormalError(c, "update error")
		return
	}
	result.Ok(c)

}

// DeleteUser
// @Summary 删除用户
// @Description 删除用户
// @Tags 用户
// @param id formData int false "id"
// @Success 200 {string} json "{"code":"200","msg":"","data":""}"
// @Router /user-delete [post]
func DeleteUser(c *gin.Context) {
	var data model.User
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
