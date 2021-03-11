package plugins

import (
	"errors"
	"sort"

	"github.com/sirupsen/logrus"
)

// GlobalPlugins 全局插件 {'name','plugin'}
var GlobalPlugins = make(map[string]OnMessagePlugin)

// GlobalPluginIDs 排序后的全局插件ID
var GlobalPluginIDs []string

var logger = logrus.WithField("bot", "register")

//Register 注册插件
func Register(plugin OnMessagePlugin) {
	info := plugin.PluginInfo()
	if info == nil {
		panic(errors.New("not found plugin infomation"))
	}
	if info.ID == "" {
		panic(errors.New("not found plugin's ID"))

	}
	if info.Name == "" {
		panic(errors.New("not found plugin's Name"))
	}
	logger.Infof("The plugin [%s] start init...", info.Name)
	plugin.PluginInit()
	GlobalPlugins[info.ID] = plugin
	logger.Infof("The plugin [%s] has been registed", info.Name)
	GlobalPluginIDs = append(GlobalPluginIDs, info.ID)
	sort.Slice(GlobalPluginIDs, func(i, j int) bool {
		return GlobalPlugins[GlobalPluginIDs[i]].PluginInfo().SortNum < GlobalPlugins[GlobalPluginIDs[j]].PluginInfo().SortNum
	})
}
