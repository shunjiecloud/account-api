package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/shunjiecloud/account/routers/api/v1"
	"github.com/shunjiecloud/pkg/middlewares"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(middlewares.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/account/v1/")
	apiv1.POST("users", v1.AddUser)

	return r
}
