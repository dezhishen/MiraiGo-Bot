package plugins

import (
	"errors"
	"sort"

	"github.com/Logiase/MiraiGo-Template/bot"
	"github.com/sirupsen/logrus"
)

var logger = logrus.WithField("bot", "register")

//RegisterOnMessagePlugin 注册消息插件
func RegisterOnMessagePlugin(plugin OnMessagePlugin) {
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
	GlobalOnMessagePlugins[info.ID] = plugin
	logger.Infof("The plugin [%s] has been registed", info.Name)
	GlobalOnMessagePluginIDs = append(GlobalOnMessagePluginIDs, info.ID)
	sort.Slice(GlobalOnMessagePluginIDs, func(i, j int) bool {
		return GlobalOnMessagePlugins[GlobalOnMessagePluginIDs[i]].PluginInfo().SortNum < GlobalOnMessagePlugins[GlobalOnMessagePluginIDs[j]].PluginInfo().SortNum
	})
}

//RegisterSchedulerPlugin 注册消息插件
func RegisterSchedulerPlugin(plugin SchedulerPlugin) {
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
	cron := plugin.Cron()
	if cron == "" {
		panic(errors.New("not found plugin's cron"))

	}
	logger.Infof("The plugin [%s] start init...", info.Name)
	plugin.PluginInit()
	err := crons.AddFunc(cron, func() {
		c := bot.Instance.QQClient
		plugin.Run(c)
	})
	if err != nil {
		panic(err)
	}
	logger.Infof("The plugin [%s] has been registed", info.Name)
}
