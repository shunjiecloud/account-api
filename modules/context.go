package modules

import (
	"fmt"

	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2/config"
)

type moduleWrapper struct {
	Redis     *redis.Client
	DefaultDB *gorm.DB
}

//ModuleContext 模块上下文
var ModuleContext moduleWrapper

//Setup 初始化Modules
func Setup() {
	//  redis
	var rConfig RedisConfig
	var err error
	if err = config.Get("config", "redis").Scan(&rConfig); err != nil {
		panic(err)
	}
	ModuleContext.Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", rConfig.Address, rConfig.Port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err = ModuleContext.Redis.Ping().Result()
	if err != nil {
		panic(err)
	}

	//  db
	var dbConfig DBConfig
	if err := config.Get("config", "db").Scan(&dbConfig); err != nil {
		panic(err)
	}
	db, err := gorm.Open(dbConfig.DbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Name))
	if err != nil {
		panic(fmt.Sprintf("%v", err.Error()))
	}
	//  连接数配置
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	//  设置表名属性
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return dbConfig.TablePrefix + defaultTableName
	}
	db.SingularTable(true)
	ModuleContext.DefaultDB = db
}
