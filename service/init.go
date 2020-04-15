package service

import (
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"scarlet/tool"
	"time"
)

var (
	redisPool        = newPool()
	mysqlClient, err = gorm.Open("mysql", "scarlet:scarlet@(127.0.0.1:3306)/scarlet?charset=utf8&parseTime=True&loc=Local")
	serverService    = NewServerService()
)

func newPool() *redis.Pool {
	return &redis.Pool{
		Dial: func() (c redis.Conn, err error) {
			return redis.Dial("tcp", "localhost:6379", redis.DialPassword("123456"))
		},
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
	}
}

func init() {
	if err != nil {
		tool.GetLogger().Fatal(err)
	}
	mysqlClient.SetLogger(tool.GetLogger().WithField("database", "mysql"))
}
