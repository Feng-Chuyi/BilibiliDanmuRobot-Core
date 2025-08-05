package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"time"

	"github.com/xbclub/BilibiliDanmuRobot-Core/config"
	"github.com/xbclub/BilibiliDanmuRobot-Core/logic"
	"github.com/xbclub/BilibiliDanmuRobot-Core/svc"
)

func main() {
	fmt.Println("====== 测试：模拟 ENTRY_EFFECT 进房欢迎逻辑（修复版） ======")

	cfg := config.Config{
		RobotName:    "小橘",
		DanmuLen:     30,
		EntryEffect:  true,
		InteractWord: true,
	}
	svcCtx := svc.NewTestServiceContext(cfg)

	// 启动弹幕处理协程
	ctx := context.Background()
	go logic.StartSendBullet(ctx, svcCtx)

	// 手动模拟 ENTRY_EFFECT 的 JSON 消息
	payload := `{
		"data": {
			"uid": 123456,
			"uinfo": {
				"base": { "name": "测试用户" },
				"guard": { "level": 0 },
				"wealth": { "level": 0 }
			}
		}
	}`

	type EntryEffectText struct {
		Data struct {
			Uid   int64 `json:"uid"`
			Uinfo struct {
				Base struct {
					Name string `json:"name"`
				} `json:"base"`
				Guard struct {
					Level int `json:"level"`
				} `json:"guard"`
				Wealth struct {
					Level int `json:"level"`
				} `json:"wealth"`
			} `json:"uinfo"`
		} `json:"data"`
	}

	var entry EntryEffectText
	_ = json.Unmarshal([]byte(payload), &entry)

	// 根据解析结果生成欢迎词（简化模拟逻辑）
	username := entry.Data.Uinfo.Base.Name
	welcome := fmt.Sprintf("欢迎 %s 来到直播间 👏", username)

	// 推送弹幕
	logic.PushToBulletSender(welcome)

	time.Sleep(2 * time.Second)
}
