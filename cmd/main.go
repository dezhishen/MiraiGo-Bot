package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/Logiase/MiraiGo-Template/bot"
	"github.com/Logiase/MiraiGo-Template/config"
	"github.com/Logiase/MiraiGo-Template/utils"

	_ "github.com/dezhiShen/MiraiGo-Bot/pkg/customplugins"
	_ "github.com/dezhiShen/MiraiGo-Bot/pkg/modules"
)

func init() {
	// migrations.Sync("file://./migrations")
	utils.WriteLogToFS()
	exists, _ := pathExists("./application.yaml")
	if !exists {
		//输入账号

		fmt.Println("请输入账号:")
		var account string
		fmt.Scanln(&account)
		//输入密码
		fmt.Println("请输入密码:")
		var password string
		fmt.Scanln(&password)
		f, err := os.Create("./application.yaml")
		defer f.Close()
		if err != nil {
			panic(err.Error())
		} else {
			_, err = f.WriteString(fmt.Sprintf("bot:\n  account: %v\n  password: %v", account, password))
			if err != nil {
				panic(err.Error())
			}
		}
		f.Close()

	}
	config.Init()
	exists, _ = pathExists("./device.json")
	if !exists {
		bot.GenRandomDevice()
	}
}
func main() {
	// 快速初始化
	bot.Init()
	// 初始化 Modules
	bot.StartService()
	// 使用协议
	// 不同协议可能会有部分功能无法使用
	// 在登陆前切换协议
	bot.UseProtocol(bot.AndroidPhone)
	// 登录
	bot.Login()
	// 刷新好友列表，群列表
	bot.RefreshList()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, os.Kill)
	<-ch
	bot.Stop()
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
