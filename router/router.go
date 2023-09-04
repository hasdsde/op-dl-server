package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "hasdsd.cn/op-dl-server/docs"
	"hasdsd.cn/op-dl-server/service"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method

		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, x-token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
	}
}
func Router() *gin.Engine {
	r := gin.Default()
	err := r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		fmt.Print("proxy error:" + err.Error())
	}
	r.Use(Cors())

	//swagger
	ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1),
	)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//ping
	r.GET("/ping", service.Ping)

	//version 版本
	r.GET("/version", service.GetVersions)
	r.PUT("/version", service.UpdateVersion)
	r.POST("/version", service.CreateVersion)
	r.POST("/version-delete", service.DeleteVersion)
	r.GET("/version-with-tag", service.GetVersionWithTag)
	r.GET("/current-version-with-tag", service.GetCurrentVersionWithTag)

	r.GET("/version-tag", service.GetVersionTag)
	r.POST("/version-tag", service.CreateVersionTag)
	r.POST("/version-tag-delete", service.DeleteVersionTag)

	//versionEvent 版本活动
	r.GET("/version-event", service.GetVersionEvent)
	r.POST("/version-event", service.CreateVersionEvent)
	r.PUT("/version-event", service.UpdateVersionEvent)
	r.POST("/version-event-delete", service.DeleteVersionEvent)
	r.GET("/version-event-with-tag", service.GetVersionEventWithTag)
	r.GET("/current-version-event-with-tag", service.GetCurrentVersionEventWithTag)

	r.GET("/version-event-tag", service.GetVersionEventTag)
	r.POST("/version-event-tag", service.CreateVersionEventTag)
	r.POST("/version-event-tag-delete", service.DeleteVersionEventTag)

	//event 限时活动
	r.GET("/event", service.GetEvent)
	r.PUT("/event", service.UpdateEvent)
	r.POST("/event", service.CreateEvent)
	r.POST("/event-delete", service.DeleteEvent)
	r.GET("/event-with-tag-detail", service.GetEventWithTagWithEventDetail)
	r.GET("/event-with-tag", service.GetEventWithTag)
	r.GET("/current-event-with-tag-detail", service.GetCurrentEventWithTagWithEventDetail)

	r.GET("/event-tag", service.GetEventTag)
	r.POST("/event-tag", service.CreateEventTag)
	r.POST("/event-tag-delete", service.DeleteEventTag)

	//目前用不上
	//user
	//r.GET("/user", service.GetUser)
	//r.POST("/user", service.CreateUser)
	//r.PUT("/user", service.UpdateUser)
	//r.POST("/user-delete", service.DeleteUser)

	r.POST("/register", service.Register)
	r.POST("/login", service.Login)

	//userVersionEvent 用户版本活动
	r.GET("/user-version-event", service.GetUserVersionEvent)
	r.POST("/user-version-event", service.CreateUserVersionEvent)
	r.PUT("/user-version-event", service.UpdateUserVersionEvent)
	r.POST("/user-version-event-delete", service.DeleteUserVersionEvent)

	//userEvent 用户活动 userEvent
	r.GET("/user-event", service.GetUserEvent)
	r.POST("/user-event", service.CreateUserEvent)
	r.PUT("/user-event", service.UpdateUserEvent)
	r.POST("/user-event-delete", service.DeleteUserEvent)

	//用户自定义 userPrivate
	r.GET("/user-private", service.GetUserPrivate)
	r.POST("/user-private", service.CreateUserPrivate)
	r.PUT("/user-private", service.UpdateUserPrivate)
	r.POST("/user-private-delete", service.DeleteUserPrivate)

	//用户每日任务
	r.GET("/daily", service.GetUserDaily)
	r.POST("/daily", service.FreshUserDaily)

	//Pool 卡池
	r.GET("/pool", service.GetPool)
	r.POST("/pool", service.UpdatePool)
	r.PUT("/pool", service.UpdatePool)
	r.POST("/pool-delete", service.DelPool)
	r.GET("/pool-with-tag", service.GetPoolWithTag)
	r.GET("/current-pool-with-tag", service.GetCurrentWithPoolTag)

	r.GET("/pool-tag", service.GetPoolTag)
	r.POST("/pool-tag", service.CreatePoolTag)
	r.POST("/pool-tag-delete", service.DeletePoolTag)

	//活动详情 eventDetail
	r.GET("/eventDetail", service.GetEventDetail)
	r.POST("/eventDetail", service.CreateEventDetail)
	r.PUT("/eventDetail", service.UpdateEventDetail)
	r.POST("/eventDetail-delete", service.DelEventDetail)

	//标签 Tag
	r.GET("/tag", service.GetTag)
	r.POST("/tag", service.CreateTag)
	r.PUT("/tag", service.UpdateTag)
	r.POST("/tag-delete", service.DeleteTag)
	return r
}
