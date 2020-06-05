package modules

import (
	"os"

	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/file"
)

type DBConfig struct {
	Host        string `json:"host"`
	User        string `json:"user"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	DbType      string `json:"dbType"`
	TablePrefix string `json:"tbPrefix"`
}

type SessionConfig struct {
	CookieName      string `json:"cookieName"`
	EnableSetCookie bool   `json:"enableSetCookie"`
	Gclifetime      int64  `json:"gclifetime"`
	Maxlifetime     int64  `json:"maxlifetime"`
	Secure          bool   `json:"secure"`
	CookieLifeTime  int    `json:"cookieLifeTime"`
	ProviderConfig  string `json:"providerConfig"`
	Domain          string `json:"domain"`
}

func setupConfig() {
	configFilePath := os.Getenv("CONFIG_FILE_PATH")
	if len(configFilePath) == 0 {
		panic("CONFIG_FILE_PATH is error")
	}
	if err := config.Load(file.NewSource(
		file.WithPath(configFilePath),
	)); err != nil {
		panic(err)
	}
}

func init() {
	setupConfig()
}
