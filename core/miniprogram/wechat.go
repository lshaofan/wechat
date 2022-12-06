package miniprogram

import (
	"github.com/lshaofan/wechat/core/credential"
	"github.com/lshaofan/wechat/core/miniprogram/config"
	"github.com/lshaofan/wechat/core/miniprogram/context"
)

// MiniProgram 微信小程序相关API
type MiniProgram struct {
	ctx *context.Context
}

func GetMiniProgram(cfg *config.Config) *MiniProgram {
	return NewMiniProgram(cfg)
}

func NewMiniProgram(cfg *config.Config) *MiniProgram {
	defaultAkHandle := credential.NewDefaultAccessToken(cfg.AppId, cfg.AppSecret, cfg.Cache)
	ctx := &context.Context{
		IAccessToken: defaultAkHandle,
		Config:       cfg,
	}
	return &MiniProgram{ctx}
}
