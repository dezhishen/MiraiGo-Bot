package plugins

import (
	"errors"

	"github.com/sirupsen/logrus"
)

// GlobalPlugins 全局插件 {'name','plugin'}
var GlobalPlugins = make(map[string]Plugin)

var logger = logrus.WithField("bot", "register")

//Register 注册插件
func Register(plugin Plugin) {
	info := plugin.PluginInfo()
	if info == nil {
		panic(errors.New("not found plugin infomation"))
	}
	logger.Infof("The plugin [%s] start init...", info.Name)
	plugin.PluginInit()
	GlobalPlugins[info.Name] = plugin
	logger.Infof("The plugin [%s] has been registed", info.Name)
}
