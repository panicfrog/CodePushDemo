package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	AccessKey string
	SecretKey string
}
