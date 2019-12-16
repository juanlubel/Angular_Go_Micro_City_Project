//+build wireinject

package main

import (
	forums "Go_Gingonic_Redis/forum"
	"Go_Gingonic_Redis/redis_server"

	"github.com/go-redis/redis"
	"github.com/google/wire"
)

func initTopicsAPI(redis *redis.Client, redisApi *redis_server.MainAPi) forums.CacheServerAPI {
	wire.Build(forums.ProvideCacheAPI)

	return forums.CacheServerAPI{}
}
