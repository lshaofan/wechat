package context

import (
	"github.com/lshaofan/wechat/core/credential"
	"github.com/lshaofan/wechat/core/miniprogram/config"
)

// Context struct
type Context struct {
	*config.Config
	credential.IAccessToken
}
