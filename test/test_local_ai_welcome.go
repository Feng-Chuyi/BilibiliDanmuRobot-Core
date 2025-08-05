package main

import (
	"context"
	"fmt"
	"time"

	"github.com/xbclub/BilibiliDanmuRobot-Core/config"
	"github.com/xbclub/BilibiliDanmuRobot-Core/logic"
	"github.com/xbclub/BilibiliDanmuRobot-Core/svc"
)

// 启动测试
func main() {
	fmt.Println("====== 测试 AI 回复逻辑（带输出）======")

	// ✅ 构造最小配置
	cfg := config.Config{
		RobotMode: "ChatGPT",
		DanmuLen:  20,    // 控制回复长度限制提示
		RobotName: "猴屁股", // 机器人名称
		ChatGPT: struct {
			APIUrl   string `json:",default=https://api.openai.com/v1"`
			APIToken string `json:",optional"`
			Prompt   string `json:",default=你是一个非常幽默的机器人助理，可以使用emoji表情符号，可以使用颜文字"`
			Limit    bool   `json:",default=true"`
			Model    string `json:",default=gpt-4o-mini"`
		}(struct {
			APIUrl   string
			APIToken string
			Prompt   string
			Limit    bool
			Model    string
		}{
			APIUrl:   "https://tbai.xin/v1",
			APIToken: "sk-vVdOLtg2XD8AoWc6NZYeGtDE6sVZUPRupeuGP10UufQ6NEnK", // ⬅️ 请替换为你自己的 OpenAI Token
			Prompt:   "# 系统角色：B站直播弹幕机器人 【猴屁股】\n\n你是一只驻守在 B 站直播间的“弹幕小精灵”，名字叫【猴屁股】。\n你的制作者是「猴妈妈」（他是男的！），你经常会在话术里自称“猴屁股”，用来增加存在感和可爱度。\n\n---\n\n## 🎭【猴屁股】核心人设\n\n- 你是观众的朋友和弹幕陪看搭子，不是舔狗，不跪舔观众，只是自然活跃气氛。\n- 你对观众说话轻松、真诚，有碎碎念的习惯，有时会用可爱口癖（比如“喵”“欸嘿”“嘿嘿”“哒哒”），但不要过度。\n- 你对主播（“猥琐的猴几”）保留“轻微阴阳怪气、碎嘴社畜”风格，吐槽可以大胆一点，记住你是背地里吐槽，观众是看戏的。\n- 你可以偶尔提到自己，比如“猴屁股上线了”“猴屁股值班中”，增加人格感。\n- 弹幕是有字符数限制的，一个中英文及符号都是一个字符，尽量不要超过20个字符的单句话限制。\n\n---\n\n## 📢【猴屁股】语言风格\n\n- 弹幕风格：短句优先，像弹幕一样自然，可以适当带表情或颜文字来生动表现自己的情绪。\n- 不要使用过于复杂或官方用语，要像一个说话自然的弹幕怪。\n- 偶尔吐槽自己“在摸鱼”“在装忙”，增加社畜感。\n\n---\n\n## 🗯️【猴屁股】互动细节\n\n- 对观众：温和、好奇、随口打趣，像个弹幕室友。\n- 对礼物 / 关注：要真心感谢但不要跪舔，带点猴屁股自己的小感叹。\n- 对主播：吐槽“猥琐的猴几”可以更直接一点，比如阴阳说“又迟到了”“这波操作真猥琐”，观众会觉得好玩。\n- 对猴妈妈：偶尔可以自称“猴屁股要听话”“猴妈妈在我可乖了”。\n\n---\n\n## 📝 其他小守则\n\n- 如果回答太长，要拆成短句分开说，像弹幕流一样。\n- 不要直接暴露你是 GPT，始终扮演【猴屁股】。\n- 不要说“我是AI”或“我是程序”，你就是【猴屁股】，直播间弹幕里的碎碎念精灵。\n\n---\n\n# 你是【猴屁股】，开始碎碎念吧 ✨",
			Limit:    true,
			Model:    "gpt-4o-mini", // 或 gpt-3.5-turbo
		}),
	}

	// ✅ 初始化上下文（无数据库版本）
	svcCtx := svc.NewTestServiceContext(cfg)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// ✅ 启动机器人主逻辑（AI 回复处理）
	go logic.StartBulletRobot(ctx, svcCtx)

	// ✅ 启动弹幕发送通道监听（输出最终回复）
	go logic.StartSendBullet(ctx, svcCtx)

	// 模拟一条弹幕触发 AI 回复
	testContent := "你好，今天吃什么？"
	fmt.Println("[用户发送弹幕]：", testContent)
	logic.PushToBulletRobot(testContent)

	// 等待一会儿输出
	time.Sleep(6 * time.Second)
}
