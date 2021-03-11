package plugins

import (
	"github.com/Logiase/MiraiGo-Template/bot"
)

type plugin interface {
	//SortNum 插件排序
	SortNum() int8
	//PluginInfo 获取插件的信息
	PluginInfo() *PluginInfo
	//PluginInit 插件初始化
	PluginInit()
}

//OnMessagePlugin 监听消息插件
type OnMessagePlugin interface {
	plugin
	//IsFireEvent 是否触发事件
	IsFireEvent(msg *MessageRequest) bool
	//IsFireNextEvent 是否触发后续事件
	IsFireNextEvent(msg *MessageRequest) bool
	//OnMessage 监听消息
	OnMessageEvent(msg *MessageRequest) (*MessageResponse, error)
}

//SchedulerPlugin 定时器插件
type SchedulerPlugin interface {
	plugin
	//Cron 表达式
	Cron() string
	//Run 要执行的方法
	Run(c *bot.Bot) error
}
