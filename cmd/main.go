package main

import (
	_ "github.com/dezhiShen/MiraiGo-Bot-Plugins/pkg/plugins"
	_ "github.com/dezhiShen/MiraiGo-Bot-Plugins/pkg/plugins/tips"
	"github.com/dezhiShen/MiraiGo-Bot/pkg/server"
)

func main() {
	server.Start()
}
