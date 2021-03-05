package module

import (
	"sync"

	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"google.golang.org/appengine/log"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/Logiase/MiraiGo-Template/bot"
	"github.com/Logiase/MiraiGo-Template/utils"
	"github.com/golang-migrate/migrate/v4"g
	log "github.com/sirupsen/logrus"
)

func init() {
	bot.RegisterModule(instance)
}

var instance = &ar{}
var logger = utils.GetModuleLogger("logiase.autoreply")
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
	db, err := sql.Open("sqlite", "/mirai/qq.db")
	if err != nil {
		log.Errorf("rotatelogs init err: %v", err)
		panic(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Errorf("rotatelogs init err: %v", err)
		panic(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"sqlite", driver)
	if err != nil {
		log.Errorf("rotatelogs init err: %v", err)
		panic(err)
	}
}

func (a *ar) PostInit() {
}

func (a *ar) Serve(b *bot.Bot) {
	b.OnGroupMessage(func(c *client.QQClient, msg *message.GroupMessage) {
		//获取群号,获取回复内容,回复
		out := autoreply(msg.ToString())
		if out == "" {
			return
		}
		m := message.NewSendingMessage().Append(message.NewText(out))
		c.SendGroupMessage(msg.GroupCode, m)
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
