package config

import (
	"gim/pkg/logger"

	"go.uber.org/zap"
)

func initLocalConf() {
	Config = config{
		MySQL:         "rtrs:Rtrs@2018++@tcp(10.192.9.23:11003)/gim?charset=utf8&parseTime=true",
		RedisIP:       "game-redis-pre-pre0.redis.dba.dandy.lan.:11282,game-redis-pre-pre1.redis.dba.dandy.lan.:11283,game-redis-pre-pre2.redis.dba.dandy.lan.:11284",
		RedisPassword: "alber123456",

		ConnectRPCAddr:  "addrs:///127.0.0.1:50000",
		LogicRPCAddr:    "addrs:///127.0.0.1:50100",
		BusinessRPCAddr: "addrs:///127.0.0.1:50200",

		TCPListenAddr:        ":8080",
		WSListenAddr:         ":8081",
		ConnectRPCListenAddr: ":50000",
		ConnectLocalAddr:     "127.0.0.1:50000",
		PushRoomSubscribeNum: 100,
		PushAllSubscribeNum:  100,

		LogicRPCListenAddr: ":50100",

		BusinessRPCListenAddr: ":50200",
	}

	logger.Level = zap.DebugLevel
	logger.Target = logger.Console
}
