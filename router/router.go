package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "hasdsd.cn/op-dl-server/docs"
	"hasdsd.cn/op-dl-server/service"
)

func Router() *gin.Engine {
	r := gin.Default()

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
	r.POST("/event-with-tag", service.CreateEventTag)
	r.POST("/event-with-tag-delete", service.DeleteEventTag)

	//user
	r.GET("/user", service.GetUser)
	r.POST("/user", service.CreateUser)
	r.PUT("/user", service.UpdateUser)
	r.POST("/user-delete", service.DeleteUser)

	r.POST("/register", service.Register)
	r.POST("/login", service.Login)
	return r
}
