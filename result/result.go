package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var SuccessCode = 200
var SuccessMsg = "操作成功"

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Result(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		code,
		msg,
		data,
	})
}
func Error(code int, msg string) Response {
	return Response{
		code,
		msg,
		nil,
	}
}

func Ok(c *gin.Context) {
	OkWithData(c, nil)
}

func OkWithData(c *gin.Context, data interface{}) {
	Result(c, SuccessCode, SuccessMsg, data)
}

func OkWithMsg(c *gin.Context, msg string) {
	Result(c, SuccessCode, msg, nil)
}

func Fail(c *gin.Context, err Response) {
	Result(c, err.Code, err.Msg, nil)
}

func FailWithMsg(c *gin.Context, code int, msg string) {
	Result(c, code, msg, nil)
}

// FailIllegalParameter 参数非法
func FailIllegalParameter(c *gin.Context) {
	Result(c, -1, "illegal parameter", nil)
}

func FailNormalError(c *gin.Context, msg string) {
	Result(c, -1, msg, nil)
}
