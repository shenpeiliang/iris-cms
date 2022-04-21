package service

import (
	"time"

	"github.com/kataras/iris/v12/sessions/sessiondb/redis"
)

var (
	//驱动
	RedisDriver = redis.Redigo()

	//数据库
	RedisDB = redis.New(redis.Config{
		Network:   "tcp",
		Addr:      "127.0.0.1:6379",
		Timeout:   time.Duration(30) * time.Second,
		MaxActive: 10,
		Password:  "",
		Database:  "",
		Prefix:    "",
		Delim:     "-",
		Driver:    RedisDriver,
	})
)
