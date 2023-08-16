package util

import (
	"crypto/md5"
	"fmt"
	"github.com/Huiyicc/mhyapi/genshin"
	"github.com/Huiyicc/mhyapi/mhyapp"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"hasdsd.cn/op-dl-server/define"
	"log"
	"strconv"
	"time"
)

// GetPageInfo 获取分页参数
func GetPageInfo(ctx *gin.Context) (int, int) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		log.Println("get param error", err.Error())
	}
	size, err := strconv.Atoi(ctx.DefaultQuery("size", define.DefaultSize))
	if err != nil {
		log.Println("get param error", err.Error())
	}
	page = (page - 1) * size
	return page, size
}

// TimeFormat 字符串转时间
func TimeFormat(str string) time.Time {
	t, err := time.Parse(time.DateTime, str)
	if err != nil {
		panic(err)
	}
	return t
}

// GenMd5 生成MD5
func GenMd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// GenerateToken 生成Token
func GenerateToken(userId int64, userName string, password string) (string, error) {
	claims := jwt.MapClaims{
		"userId":   userId,
		"userName": userName,
		"exp":      define.TokenExp.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(password + TokenSecret))
}

// ParseToken 解析Token
func ParseToken(token string, password string) (*jwt.Token, error) {
	claims := jwt.MapClaims{}
	result, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(password + TokenSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if result.Valid {
		return result, nil
	}
	return nil, jwt.ValidationError{}
}

// GetDailyStatus 根据token获取每日委托
func GetDailyStatus(token string) *genshin.NoteInfo {

	app := mhyapp.AppCore{}
	if err := app.LoginToCookies(token); err != nil {
		panic(err)
	}
	appCookiesStr := app.Cookies.GetStr()

	genshinCore, err := genshin.NewCore(appCookiesStr)
	if err != nil {
		panic(err)
	}

	// genshinCore.GameInfo = *gameInfo
	noteInfo, err := genshinCore.GetNoteInfo()
	if err != nil {
		fmt.Print(err.Error())
		return nil
	}
	return noteInfo
}
