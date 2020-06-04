package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/shunjiecloud/account-api/schemas"
	"github.com/shunjiecloud/pkg/app"
)

func CreateUser(c *gin.Context) {
	//  ctx
	appCtx := app.AppContext{
		GinCtx: c,
	}
	var request schemas.CreateUserRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		appCtx.WriteError(err)
		return
	}
	//  TODO : to account-srv
	appCtx.OkResponse()
}
