package v1

import (
	"context"

	"github.com/gin-gonic/gin"
	proto_account "github.com/shunjiecloud-proto/account/proto"
	"github.com/shunjiecloud/account-api/modules"
	"github.com/shunjiecloud/pkg/app"
)

type CreateUserRequest struct {
	CaptchaId       string `json:"captcha_id" binding:"max=32"`
	CaptchaSolution string `json:"captcha_solution" binding:"max=32"`
	Name            string `json:"name" binding:"min=3,max=16"`
	Password        string `json:"password"`
	CryptoId        int64  `json:"crypto_id" binding:"required"`
	Mail            string `json:"mail" binding:"email"`
	Phone           string `json:"phone" binding:"len=11"`
	Gender          int    `json:"gender" binding:"min=0,max=1"`
}

type CreateUserResponse struct {
	UserId int64  `json:"user_id"`
	Name   string `json:"name"`
	Mail   string `json:"mail"`
	Phone  string `json:"phone"`
	Gender int    `json:"gender"`
	Avatar string `json:"avatar"`
}

//
// @Summary 创建用户
// @tags users
// @Description 创建用户
// @Accept  json
// @Produce  json
// @Param createUserRequest {object} body CreateUserRequest true "密码先sha1，之后使用公钥加密，再stdBase64编码。"
// @Success 200 {object} CreateUserResponse
// @Router /account/v1/users [post]
func CreateUser(c *gin.Context) {
	//  ctx
	appCtx := app.AppContext{
		GinCtx: c,
	}
	var request CreateUserRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		appCtx.WriteError(err)
		return
	}
	//  调用srv注册用户
	signUpResp, err := modules.ModuleContext.AccountSrvClient.SignUp(context.Background(), &proto_account.SignUpRequest{
		Name:            request.Name,
		Password:        request.Password,
		CryptoId:        request.CryptoId,
		Mail:            request.Mail,
		Phone:           request.Phone,
		Gender:          int32(request.Gender),
		CaptchaId:       request.CaptchaId,
		CaptchaSolution: request.CaptchaSolution,
	})
	if err != nil {
		appCtx.WriteError(err)
		return
	}

	//  下发cookie
	sess, _ := modules.ModuleContext.SessionMgr.SessionStart(c.Writer, c.Request)
	defer sess.SessionRelease(c.Writer)
	sess.Set("userId", signUpResp.UserId)

	//  返回
	var ret CreateUserResponse
	ret.UserId = signUpResp.UserId
	ret.Name = signUpResp.Name
	ret.Mail = signUpResp.Mail
	ret.Phone = signUpResp.Phone
	ret.Gender = int(signUpResp.Gender)
	ret.Avatar = signUpResp.Avatar
	appCtx.WriteJson(200, &ret)
}
