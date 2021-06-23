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
		panic(errors.New("not found OnMessagePlugin's info"))
	}
	if info.ID == "" {
		panic(errors.New("not found OnMessagePlugin's ID"))

	}
	if info.Name == "" {
		panic(errors.New("not found plugin's Name"))
	}
	logger.Infof("The OnMessagePlugin [%s] start init...", info.Name)
	plugin.PluginInit()
	GlobalOnMessagePlugins[info.ID] = plugin
	logger.Infof("The OnMessagePlugin [%s] has been registed", info.Name)
	GlobalOnMessagePluginIDs = append(GlobalOnMessagePluginIDs, info.ID)
	sort.Slice(GlobalOnMessagePluginIDs, func(i, j int) bool {
		return GlobalOnMessagePlugins[GlobalOnMessagePluginIDs[i]].SortNum() < GlobalOnMessagePlugins[GlobalOnMessagePluginIDs[j]].SortNum()
	})
}

//RegisterSchedulerPlugin 注册消息插件
func RegisterSchedulerPlugin(plugin SchedulerPlugin) {
	info := plugin.PluginInfo()
	if info == nil {
		panic(errors.New("not found SchedulerPlugin's info"))
	}
	if info.ID == "" {
		panic(errors.New("not found SchedulerPlugin's ID"))

	}
	if info.Name == "" {
		panic(errors.New("not found SchedulerPlugin's Name"))
	}
	cron := plugin.Cron()
	if cron == "" {
		panic(errors.New("not found SchedulerPlugin's cron"))

	}
	logger.Infof("The SchedulerPlugin [%s] start init...", info.Name)
	plugin.PluginInit()
	err := Crons.AddFunc(cron, func() {
		plugin.Run(bot.Instance)
	})
	if err != nil {
		panic(err)
	}
	logger.Infof("The SchedulerPlugin [%s] has been registed", info.Name)
}

func RegisterCoroutinePlugin(plugin CoroutinePlugin) {
	info := plugin.PluginInfo()
	if info == nil {
		panic(errors.New("not CoroutinePlugin's info"))
	}
	if info.ID == "" {
		panic(errors.New("not found CoroutinePlugin's ID"))

	}
	if info.Name == "" {
		panic(errors.New("not found CoroutinePlugin's Name"))
	}
	plugin.PluginInit()
	logger.Infof("The CoroutinePlugin [%s] has been registed", info.Name)
	GlobalCoroutinePlugins = append(GlobalCoroutinePlugins, plugin)
	// go plugin.Run(bot.Instance)
}
