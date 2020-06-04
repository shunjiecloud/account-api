package modules

import (
	"github.com/micro/go-micro/v2"
	proto_account "github.com/shunjiecloud-proto/account/proto"
)

type moduleWrapper struct {
	AccountSrvClient proto_account.AccountService
}

//ModuleContext 模块上下文
var ModuleContext moduleWrapper

//Setup 初始化Modules
func Setup() {
	//  account-srv client
	m_service := micro.NewService()
	ModuleContext.AccountSrvClient = proto_account.NewAccountService("go.micro.srv.account", m_service.Client())
}
