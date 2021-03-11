package plugins

type plugin interface {
	//PluginInfo 获取插件的信息
	PluginInfo() *PluginInfo
	//PluginInit 插件初始化
	PluginInit()
	//IsFireEvent 是否触发事件
	IsFireEvent(msg *MessageRequest) bool
	//IsFireNextEvent 是否触发后续事件
	IsFireNextEvent(msg *MessageRequest) bool
}

//OnMessagePlugin 监听消息插件
type OnMessagePlugin interface {
	plugin
	//OnMessage 监听消息
	OnMessageEvent(msg *MessageRequest) (*MessageResponse, error)
}

//SchedulerPlugin 定时器插件
type SchedulerPlugin interface {
	plugin
	Run() error
}
