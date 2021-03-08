package modules

import (
	"sync"

	"github.com/Logiase/MiraiGo-Template/bot"
	"github.com/Logiase/MiraiGo-Template/utils"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"github.com/dezhiShen/MiraiGo-Bot/pkg/engine"
)

func init() {
	bot.RegisterModule(instance)
}

var instance = &ar{}
var logger = utils.GetModuleLogger("dezhiShen.reply")
var tem map[string]string

type ar struct {
}

func (a *ar) MiraiGoModule() bot.ModuleInfo {
	return bot.ModuleInfo{
		ID:       "dezhiShen.reply",
		Instance: instance,
	}
}

func (a *ar) Init() {
}

func (a *ar) PostInit() {
}

func (a *ar) Serve(b *bot.Bot) {
	b.OnGroupMessage(func(c *client.QQClient, msg *message.GroupMessage) {
		out := engine.ReplyToGroupMessage(msg)
		if out == "" {
			return
		}
		m := message.NewSendingMessage().Append(message.NewText(out))
		go c.SendGroupMessage(msg.GroupCode, m)
	})

	b.OnPrivateMessage(func(c *client.QQClient, msg *message.PrivateMessage) {
		out := engine.ReplyToPrivateMessage(msg)
		if out == "" {
			return
		}
		m := message.NewSendingMessage().Append(message.NewText(out))
		go c.SendPrivateMessage(msg.Sender.Uin, m)
	})
}

func (a *ar) Start(bot *bot.Bot) {
}

func (a *ar) Stop(bot *bot.Bot, wg *sync.WaitGroup) {
	defer wg.Done()
}
