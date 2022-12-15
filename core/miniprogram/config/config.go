package config

import "github.com/lshaofan/wechat/store/cache"

type Config struct {
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	Cache     cache.Operation
}

func NewConfig(appId, appSecret string, c cache.RedisConfig) *Config {
	return &Config{
		AppId:     appId,
		AppSecret: appSecret,
		Cache:     *cache.NewOperation(&c),
	}
}
