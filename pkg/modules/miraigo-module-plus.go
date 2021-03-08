package modules

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Logiase/MiraiGo-Template/bot"
	"github.com/Logiase/MiraiGo-Template/utils"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"github.com/dezhiShen/MiraiGo-Bot/pkg/entity"
	"gopkg.in/fatih/set.v0"
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
		out := replyToGroupMessage(msg)
		if out == "" {
			return
		}
		m := message.NewSendingMessage().Append(message.NewText(out))
		c.SendGroupMessage(msg.GroupCode, m)
	})

	b.OnPrivateMessage(func(c *client.QQClient, msg *message.PrivateMessage) {
		out := replyToPrivateMessage(msg)
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

// 获取回复群消息
func replyToGroupMessage(msg *message.GroupMessage) string {
	//获取回复内容
	messageText := msg.ToString()
	//判断是否为命令
	if strings.HasPrefix(messageText, ".") {
		//切割内容
		messageArray := strings.Split(messageText, " ")
		//第一个为命令
		command := messageArray[0]
		msgContext := getGroupMessageContext(msg)
		rule := getRule("group", msg.GroupCode, command)
		answer := runRule(rule, msgContext)
		msgContext["$answer"] = answer
		log.Printf("%v,%v", msgContext, rule)
		out := replaceTemplate(rule.RespTemplate, msgContext)
		return out

	}
	return ""
}

// 获取回复私聊消息
func replyToPrivateMessage(msg *message.PrivateMessage) string {
	//获取回复内容
	messageText := msg.ToString()
	//判断是否为命令
	if strings.HasPrefix(messageText, ".") {
		//切割内容
		messageArray := strings.Split(messageText, " ")
		//第一个为命令
		command := messageArray[0]
		msgContext := getPrivateMessageContext(msg)
		rule := getRule("group", msg.Sender.Uin, command)
		answer := runRule(rule, msgContext)
		msgContext["$answer"] = answer
		log.Printf("%v,%v", msgContext, rule)
		out := replaceTemplate(rule.RespTemplate, msgContext)
		return out

	}
	return ""
}

func getGroupMessageContext(msg *message.GroupMessage) map[string]string {
	result := make(map[string]string)
	messageText := msg.ToString()
	messageArray := strings.Split(messageText, " ")
	if len(messageArray) > 1 {
		params := messageArray[1:]
		for index, value := range params {
			result[fmt.Sprintf("$%v", index+1)] = value
		}
	}
	result["$nickName"] = msg.Sender.CardName
	return result
}

func getPrivateMessageContext(msg *message.PrivateMessage) map[string]string {
	result := make(map[string]string)
	messageText := msg.ToString()
	messageArray := strings.Split(messageText, " ")
	if len(messageArray) > 1 {
		params := messageArray[1:]
		for index, value := range params {
			result[fmt.Sprintf("$%v", index+1)] = value
		}
	}
	result["$nickName"] = msg.Sender.Nickname
	return result
}

func getRule(messageType string, id int64, command string) *entity.Rule {
	if command == ".r" {
		return &entity.Rule{
			Command:      ".r",
			ID:           ".r",
			Max:          100,
			Min:          0,
			Type:         "randomMath",
			RespTemplate: "[$nickName]掷出了:$answer",
		}
	}
	if command == ".rc" {
		return &entity.Rule{
			Command:      ".rc",
			ID:           ".rc",
			Max:          100,
			Min:          0,
			Type:         "randomItem",
			RespTemplate: "[$nickName]掷出了:$answer",
		}
	}
	return nil
}

func replaceTemplate(temp string, context map[string]string) string {
	if temp == "" {
		return ""
	}
	for key, exp := range context {
		//fmt.Println("exp",exp)
		exp = strings.TrimSpace(exp)
		temp = strings.Replace(temp, key, exp, -1)
	}
	return temp
}

func runRule(rule *entity.Rule, context map[string]string) string {
	if rule.Type == "randomMath" {
		rand.Seed(time.Now().UnixNano())
		v := rand.Intn(rule.Max-rule.Min) + rule.Min
		return fmt.Sprint(v)
	}
	if rule.Type == "randomItem" {
		setKey := context["$1"]
		set := getSet(setKey)
		if set == nil {
			return ""
		}
		countStr, ok := context["$2"]
		var count int
		if !ok || countStr == "" {
			count = 1
		} else {
			var err error
			count, err = strconv.Atoi(countStr)
			if err != nil {
				return "参数不合法"
			}
		}
		out := ""
		for i := 0; i < count; i++ {
			v := set.Pop()
			out += (fmt.Sprint(v) + ",")
		}
		return out[:len(out)-1]

	}

	return ""
}

// getSet 获取选项
func getSet(key string) set.Interface {
	if key == "标点符号" {
		result := set.New(set.ThreadSafe)
		result.Add("!")
		result.Add(",")
		result.Add("。")
		result.Add("...")
		return result
	}
	return nil
}
