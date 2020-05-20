package service

import (
	"github.com/gomodule/redigo/redis"
	"scarlet/common"
	"scarlet/tool"
)

type UriService struct{}

var configService = NewConfigService()

func NewURIService() *UriService {
	return new(UriService)
}

func (u *UriService) Add(uri common.URI) {
	mysqlClient.Create(&uri)
	return
}

func (u *UriService) Own(uriId uint, serverId uint) bool {
	var num int
	mysqlClient.Table("uris").Where("id = ? and server_id = ?", uriId, serverId).Count(&num)
	if num >= 1 {
		return true
	} else {
		return false
	}
}

func (u *UriService) Delete(uri common.URI) {
	mysqlClient.Delete(uri)
	return
}

func (u *UriService) Get(uriId uint) common.URI {
	var uri common.URI
	uri.ID = uriId
	mysqlClient.First(&uri)
	return uri
}

func (u *UriService) GetWithSwitch(uriId uint) common.URI {
	var uri common.URI
	uri.ID = uriId
	mysqlClient.First(&uri)
	key := tool.CustomConfigKeyGen(uri.Host, uri.Path)
	val, err := redis.Values(redisPool.Get().Do("hgetall", key))
	if err != nil {
		panic(err)
	}
	var switc common.CustomSwitch
	err = redis.ScanStruct(val, &switc)
	if err != nil {
		panic(err)
	}
	uri.Switch = switc
	return uri
}

func (u *UriService) GetByServerID(serverID uint) (uris []common.URI) {
	mysqlClient.Where("server_id = ?", serverID).Find(&uris)
	return
}

func (u *UriService) GetByServerIDWithSwitch(serverID uint) (uris []common.URI) {
	mysqlClient.Where("server_id = ?", serverID).Find(&uris)
	for i, uri := range uris {
		key := tool.CustomConfigKeyGen(uri.Host, uri.Path)
		val, err := redis.Values(redisPool.Get().Do("hgetall", key))
		if err != nil {
			panic(err)
		}
		var switc common.CustomSwitch
		err = redis.ScanStruct(val, &switc)
		if err != nil {
			panic(err)
		}
		uris[i].Switch = switc
	}
	return
}

// 通过serverId 和 path 检查是否已存在
func (u *UriService) Exist(uri common.URI) bool {
	var num int
	mysqlClient.Debug().Table("uris").Where("server_id = ? and path = ?", uri.ServerID, uri.Path).Count(&num)
	if num >= 1 {
		return true
	} else {
		return false
	}
}
