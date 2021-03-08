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
	"github.com/prometheus/common/log"
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
		//获取回复内容
		messageText := msg.ToString()
		//判断是否为命令
		if strings.HasPrefix(messageText, ".") {
			//切割内容
			messageArray := strings.Split(messageText, " ")
			//第一个为命令
			command := messageArray[0]
			params := messageArray[1:]
			log.Debug("接收到命令[%v]参数为[%v]", command, params)
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
