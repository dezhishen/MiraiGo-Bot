package engine

import (
	"fmt"
	"strings"

	"github.com/Mrs4s/MiraiGo/message"
	"github.com/dezhiShen/MiraiGo-Bot/pkg/db"
	"github.com/dezhiShen/MiraiGo-Bot/pkg/entity"
)

//ReplyToGroupMessage 获取回复群消息
func ReplyToGroupMessage(msg *message.GroupMessage) string {
	//获取回复内容
	messageText := msg.ToString()
	//判断是否为命令
	if strings.HasPrefix(messageText, ".") {
		//切割内容
		messageArray := strings.Split(messageText, " ")
		//第一个为命令
		command := messageArray[0]
		msgContext := GetGroupMessageContext(msg)
		rule := db.GetRule("group", msg.GroupCode, command)
		if rule == nil {
			return ""
		}
		answer, err := RunRule(rule, msgContext)
		if err != nil {
			return err.Error()
		}
		msgContext["$answer"] = answer
		resp := db.GetResp(rule.RespID)
		out := replaceTemplate(resp, msgContext)
		return out

	}
	return ""
}

//ReplyToPrivateMessage 获取回复私聊消息
func ReplyToPrivateMessage(msg *message.PrivateMessage) string {
	//获取回复内容
	messageText := msg.ToString()
	//判断是否为命令
	if strings.HasPrefix(messageText, ".") {
		//切割内容
		messageArray := strings.Split(messageText, " ")
		//第一个为命令
		command := messageArray[0]
		msgContext := GetPrivateMessageContext(msg)
		rule := db.GetRule("group", msg.Sender.Uin, command)
		if rule == nil {
			return ""
		}
		answer, err := RunRule(rule, msgContext)
		if err != nil {
			return err.Error()
		}
		msgContext["$answer"] = answer
		resp := db.GetResp(rule.RespID)
		out := replaceTemplate(resp, msgContext)
		return out

	}
	return ""
}

//GetGroupMessageContext GetGroupMessageContext
func GetGroupMessageContext(msg *message.GroupMessage) map[string]string {
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

//GetPrivateMessageContext GetPrivateMessageContext
func GetPrivateMessageContext(msg *message.PrivateMessage) map[string]string {
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

func replaceTemplate(resp *entity.Resp, context map[string]string) string {
	if resp == nil {
		return ""
	}
	template := resp.Template
	for key, exp := range context {
		//fmt.Println("exp",exp)
		exp = strings.TrimSpace(exp)
		template = strings.Replace(template, key, exp, -1)
	}
	return template
}
