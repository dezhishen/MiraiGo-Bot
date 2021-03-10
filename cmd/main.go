package main

import (
	_ "github.com/dezhiShen/MiraiGo-Bot/pkg/customplugins"
	"github.com/dezhiShen/MiraiGo-Bot/pkg/server"
)

func main() {
	server.Start()
}
