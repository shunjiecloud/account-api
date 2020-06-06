package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/shunjiecloud/account-api/docs"
	v1 "github.com/shunjiecloud/account-api/routers/api/v1"
	"github.com/shunjiecloud/pkg/middlewares"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(middlewares.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/account/v1/")

	url := ginSwagger.URL("https://api.shunjiecloud.com/account/v1/swagger/doc.json") // The url pointing to API definition
	apiv1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	apiv1.POST("users", v1.CreateUser)
	apiv1.GET("users", v1.GetUserInfo)

	apiv1.POST("signin", v1.SignIn)

	return r
}
