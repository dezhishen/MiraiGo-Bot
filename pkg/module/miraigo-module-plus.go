package module

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"

	"github.com/Logiase/MiraiGo-Template/bot"
	"github.com/Logiase/MiraiGo-Template/utils"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
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
		//获取群号,获取回复内容,回复
		messageText := msg.ToString()
		if strings.HasPrefix(messageText, ".") {
			messageArray := strings.Split(messageText, " ")
			command := messageArray[0]
			var out string
			if command == ".r" {
				out = fmt.Sprint(rand.Intn(100))
			}
			m := message.NewSendingMessage().Append(message.NewText(out))
			c.SendGroupMessage(msg.GroupCode, m)
		} else {
			return
		}
	})

	b.OnPrivateMessage(func(c *client.QQClient, msg *message.PrivateMessage) {
		out := autoreply(msg.ToString())
		if out == "" {
			return
		}
		m := message.NewSendingMessage().Append(message.NewText(out))
		c.SendPrivateMessage(msg.Sender.Uin, m)
	})
}

func (a *ar) Start(bot *bot.Bot) {
}

func (a *ar) Stop(bot *bot.Bot, wg *sync.WaitGroup) {
	defer wg.Done()
}

func autoreply(in string) string {
	out, ok := tem[in]
	if !ok {
		return ""
	}
	return out
}
