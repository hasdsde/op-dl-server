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

	//version
	r.GET("/version", service.GetVersions)
	r.GET("/version-with-tag", service.GetVersionWithTag)

	return r
}
