package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/shunjiecloud/account/models"
	"github.com/shunjiecloud/account/schemas"
	"github.com/shunjiecloud/pkg/app"
)

func AddUser(c *gin.Context) {
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
	user := models.User{
		Name:     request.Name,
		Password: request.Password,
		Avatar:   "",
		Gender:   request.Gender,
		Mail:     request.Mail,
		Phone:    request.Phone,
	}
	err = models.AddUser(&user)
	if err != nil {
		appCtx.WriteError(err)
		return
	}
	appCtx.OkResponse()
}
