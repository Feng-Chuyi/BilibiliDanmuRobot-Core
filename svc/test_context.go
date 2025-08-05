package svc

import (
	"github.com/xbclub/BilibiliDanmuRobot-Core/config"
)

func NewTestServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       &c,
		OtherSideUid: make(map[int64]bool),
		// 其余 DB 相关模型不初始化
	}
}
