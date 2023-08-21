package router

import (
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
	r.POST("/version-with-tag-delete", service.DeleteVersionWithTag)
	r.POST("/version-with-tag", service.CreateVersionTag)

	//versionEvent 版本活动
	r.GET("/version-event", service.GetVersionEvent)
	r.GET("/version-event-with-tag", service.GetVersionEventWithTag)
	r.POST("/version-event", service.CreateVersionEvent)
	r.PUT("/version-event", service.UpdateVersionEvent)
	r.POST("/version-event-delete", service.DeleteVersionEvent)
	r.POST("/version-event-with-tag", service.CreateVersionEventTag)
	r.POST("/version-event-with-tag-delete", service.DeleteVersionEventTag)

	//event 限时活动
	r.GET("/event", service.GetEvent)
	r.PUT("/event", service.UpdateEvent)
	r.POST("/event", service.CreateEvent)
	r.POST("/event-delete", service.DeleteEvent)
	r.GET("/event-with-tag", service.GetEventWithTag)
	r.GET("/event-with-tag-detail", service.GetEventWithTagWithEventDetail)
	r.POST("/event-with-tag", service.CreateEventTag)
	r.POST("/event-with-tag-delete", service.DeleteEventTag)

	//user
	r.GET("/user", service.GetUser)
	r.POST("/user", service.CreateUser)
	r.PUT("/user", service.UpdateUser)
	r.POST("/user-delete", service.DeleteUser)

	r.POST("/register", service.Register)
	r.POST("/login", service.Login)

	//userVersionEvent 用户版本活动
	r.GET("/user-version-event", service.GetUserVersionEvent)
	r.POST("/user-version-event", service.CreateUserVersionEvent)
	r.PUT("/user-version-event", service.UpdateUserVersionEvent)
	r.POST("/user-version-event-delete", service.DeleteUserVersionEvent)

	//userEvent 用户活动
	r.GET("/user-event", service.GetUserEvent)
	r.POST("/user-event", service.CreateUserEvent)
	r.PUT("/user-event", service.UpdateUserEvent)
	r.POST("/user-event-delete", service.DeleteUserEvent)

	//用户自定义
	r.GET("/user-private", service.GetUserPrivate)
	r.POST("/user-private", service.CreateUserPrivate)
	r.PUT("/user-private", service.UpdateUserPrivate)
	r.POST("/user-private-delete", service.DeleteUserPrivate)

	//用户每日任务
	r.GET("/daily", service.GetUserDaily)
	r.POST("/daily", service.FreshUserDaily)

	//卡池
	r.GET("/pool", service.GetPool)
	r.POST("/pool", service.UpdatePool)
	r.PUT("/pool", service.UpdatePool)
	r.POST("/pool-delete", service.DelPool)

	//活动详情
	r.GET("/eventDetail", service.GetEventDetail)
	r.POST("/eventDetail", service.CreateEventDetail)
	r.PUT("/eventDetail", service.UpdateEventDetail)
	r.POST("/eventDetail-delete", service.DelEventDetail)
	return r
}
