package v1

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	proto_account "github.com/shunjiecloud-proto/account/proto"
	"github.com/shunjiecloud/account-api/modules"
	"github.com/shunjiecloud/pkg/app"
)

type GetUserInfoRequest struct {
	UserId int64 `form:"user_id"`
}

type GetUserInfoResponse struct {
	UserId int64  `json:"user_id"`
	Name   string `json:"name"`
	Mail   string `json:"mail"`
	Phone  string `json:"phone"`
	Gender int    `json:"gender"`
	Avatar string `json:"avatar"`
}

//
// @Summary 获取用户信息
// @tags users
// @Description 获取用户信息
// @Produce  json
// @Param user_id query int false "用户id，不带userId，则获取自己的userInfo"
// @Success 200 {object} GetUserInfoResponse
// @Router /account/v1/users [get]
func GetUserInfo(c *gin.Context) {
	//  ctx
	appCtx := app.AppContext{
		GinCtx: c,
	}
	var request GetUserInfoRequest
	err := c.ShouldBindQuery(&request)
	if err != nil {
		appCtx.WriteError(err)
		return
	}
	//  TODO : 检查是否登录
	sess, _ := modules.ModuleContext.SessionMgr.SessionStart(c.Writer, c.Request)
	defer sess.SessionRelease(c.Writer)
	userIdInf := sess.Get("userId")
	if userIdInf == nil {
		appCtx.WriteError(errors.New("用户未登录"))
		return
	}

	//  不带userId，则获取自己的userInfo
	if request.UserId == 0 {
		request.UserId = userIdInf.(int64)
	}

	//  调用srv注册用户
	getUserInfoResp, err := modules.ModuleContext.AccountSrvClient.UserInfo(context.Background(), &proto_account.UserInfoRequest{
		UserId: request.UserId,
	})
	if err != nil {
		appCtx.WriteError(err)
		return
	}

	//  返回
	var ret GetUserInfoResponse
	ret.UserId = getUserInfoResp.UserId
	ret.Name = getUserInfoResp.Name
	ret.Mail = getUserInfoResp.Mail
	ret.Phone = getUserInfoResp.Phone
	ret.Gender = int(getUserInfoResp.Gender)
	ret.Avatar = getUserInfoResp.Avatar
	appCtx.WriteJson(200, &ret)
}
