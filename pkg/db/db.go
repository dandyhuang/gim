package db

import (
	"fmt"
	"gim/config"
	"gim/pkg/logger"
	"gim/pkg/util"
	"strings"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB        *gorm.DB
	RedisCli  *redis.ClusterClient
	RedisUtil *util.RedisUtil
)

func init() {
	InitMysql(config.Config.MySQL)
	InitRedis(config.Config.RedisIP, config.Config.RedisPassword)
}

// InitMysql 初始化MySQL
func InitMysql(dataSource string) {
	logger.Logger.Info("init mysql")
	var err error
	DB, err = gorm.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	DB.SingularTable(true)
	DB.LogMode(true)
	logger.Logger.Info("init mysql ok")
}

// InitRedis 初始化Redis
func InitRedis(addr, password string) {
	address :=strings.Split(addr, ",")
	logger.Logger.Info("init redis")
	RedisCli := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: address,
	})
	//RedisCli = redis.NewClient(&redis.Options{
	//	Addr:     addr,
	//	DB:       0,
	//	Password: password,
	//})

	_, err := RedisCli.Ping().Result()
	if err != nil {
		panic(err)
	}

	RedisUtil = util.NewRedisUtil(RedisCli)
	logger.Logger.Info("init redis ok")
}

// InitByTest 初始化数据库配置，仅用在单元测试
func InitByTest() {
	fmt.Println("init db")
	logger.Target = logger.Console

	InitMysql(config.Config.MySQL)
	InitRedis(config.Config.RedisIP, config.Config.RedisPassword)
}
