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
	//版本活动
	r.GET("/version-event", service.GetVersionEvent)
	r.GET("/version-event-with-tag", service.GetVersionEventWithTag)
	return r
}
