package plugins

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
