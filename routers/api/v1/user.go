package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/shunjiecloud/account-api/modules"
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
	sess, _ := modules.ModuleContext.SessionMgr.SessionStart(c.Writer, c.Request)
	defer sess.SessionRelease(c.Writer)
	appCtx.OkResponse()
}
