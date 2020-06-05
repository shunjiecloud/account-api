package modules

import (
	"fmt"

	"github.com/astaxie/beego/session"
	_ "github.com/astaxie/beego/session/mysql"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/config"
	proto_account "github.com/shunjiecloud-proto/account/proto"
)

type moduleWrapper struct {
	AccountSrvClient proto_account.AccountService
	SessionMgr       *session.Manager
}

//ModuleContext 模块上下文
var ModuleContext moduleWrapper

//Setup 初始化Modules
func Setup() {
	//  account-srv client
	m_service := micro.NewService()
	ModuleContext.AccountSrvClient = proto_account.NewAccountService("go.micro.srv.account", m_service.Client())

	//  session
	var dbConfig DBConfig
	var sessionConfig SessionConfig
	if err := config.Get("config", "db").Scan(&dbConfig); err != nil {
		panic(err)
	}
	if err := config.Get("config", "session").Scan(&sessionConfig); err != nil {
		panic(err)
	}

	providerConfig := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Name)
	managerConfig := session.ManagerConfig{
		CookieName:      sessionConfig.CookieName,
		EnableSetCookie: sessionConfig.EnableSetCookie,
		Gclifetime:      sessionConfig.Gclifetime,
		Maxlifetime:     sessionConfig.Maxlifetime,
		Secure:          sessionConfig.Secure,
		CookieLifeTime:  sessionConfig.CookieLifeTime,
		ProviderConfig:  providerConfig,
		Domain:          sessionConfig.Domain,
	}
	var err error
	ModuleContext.SessionMgr, err = session.NewManager("mysql", &managerConfig)
	if err != nil {
		panic(err)
	}
	go ModuleContext.SessionMgr.GC()
}
