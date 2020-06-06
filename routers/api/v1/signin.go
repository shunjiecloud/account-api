package v1

import (
	"context"

	"github.com/gin-gonic/gin"
	proto_account "github.com/shunjiecloud-proto/account/proto"
	"github.com/shunjiecloud/account-api/modules"
	"github.com/shunjiecloud/pkg/app"
)

type SignInRequest struct {
	// CaptchaId       string `json:"captcha_id" binding:"max=32,required"`
	// CaptchaSolution string `json:"captcha_solution" binding:"max=32,required"`
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignInResponse struct {
	UserId int64  `json:"user_id"`
	Name   string `json:"name"`
	Mail   string `json:"mail"`
	Phone  string `json:"phone"`
	Gender int    `json:"gender"`
	Avatar string `json:"avatar"`
}

//
// @Summary 用户登录
// @tags users
// @Description 用户登录
// @Accept  json
// @Produce  json
// @Param SignInRequest {object} body SignInRequest true "密码先sha1，之后使用公钥加密，再stdBase64编码。"
// @Success 200 {object} SignInResponse
// @Router /account/v1/signin [post]
func SignIn(c *gin.Context) {
	//  ctx
	appCtx := app.AppContext{
		GinCtx: c,
	}
	var request SignInRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		appCtx.WriteError(err)
		return
	}
	//  调用srv登录用户
	signInResp, err := modules.ModuleContext.AccountSrvClient.SignIn(context.Background(), &proto_account.SignInRequest{
		Account:  request.Account,
		Password: request.Password,
	})
	if err != nil {
		appCtx.WriteError(err)
		return
	}

	//  下发cookie
	sess, _ := modules.ModuleContext.SessionMgr.SessionStart(c.Writer, c.Request)
	defer sess.SessionRelease(c.Writer)
	sess.Set("userId", signInResp.UserId)

	//  返回
	var ret SignInResponse
	ret.UserId = signInResp.UserId
	ret.Name = signInResp.Name
	ret.Mail = signInResp.Mail
	ret.Phone = signInResp.Phone
	ret.Gender = int(signInResp.Gender)
	ret.Avatar = signInResp.Avatar
	appCtx.WriteJson(200, &ret)
}
