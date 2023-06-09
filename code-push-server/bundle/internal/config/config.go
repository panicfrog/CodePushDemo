package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	AccessKey string
	SecretKey string
	Mysql     struct {
		DataSource string
	}
	CacheRedis cache.CacheConf
}
