package plugins

//Plugin 插件
type Plugin interface {
	//PluginInfo 获取插件的信息
	PluginInfo() *PluginInfo
	//PluginInit 插件初始化
	PluginInit()
	//IsFireEvent 是否触发事件
	IsFireEvent(msg *MessageRequest) bool
	//OnMessage 监听消息
	OnMessageEvent(msg *MessageRequest) (*MessageResponse, error)
	//IsFireNextEvent 是否触发后续事件
	IsFireNextEvent(msg *MessageRequest) bool
}

// NoInitPlugin 无需初始化
type NoInitPlugin struct {
}

// PluginInit 简单插件初始化
func (p NoInitPlugin) PluginInit() {

}

// AlwaysFireEventPlugin 总是触发的插件
type AlwaysFireEventPlugin struct {
}

// IsFireEvent 简单插件初始化
func (p AlwaysFireEventPlugin) IsFireEvent(msg *MessageRequest) bool {
	return true

}

// AlwaysFireNextEventPlugin 总是触发下一个插件的插件
type AlwaysFireNextEventPlugin struct {
}

// IsFireNextEvent IsFireNextEvent
func (p AlwaysFireNextEventPlugin) IsFireNextEvent(msg *MessageRequest) bool {
	return true

}

// AlwaysNotFireNextEventPlugin 总是不触发下一个插件的插件
type AlwaysNotFireNextEventPlugin struct {
}

// IsFireNextEvent IsFireNextEvent
func (p AlwaysNotFireNextEventPlugin) IsFireNextEvent(msg *MessageRequest) bool {
	return false

}
